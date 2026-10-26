+++
title = "大模型API的协议转换实践"
description = "从 OpenAI 与 Anthropic 的消息结构出发，理解工具调用、SSE 事件、会话状态和能力边界如何影响协议转换。"
date = '2026-10-26T00:00:00+08:00'
draft = false
tags = ["LLM", "协议转换", "OpenAI", "Anthropic", "SSE"]
+++

接入多个大模型服务时，经常会遇到这样的需求：客户端已经使用一种 API，希望切换模型服务后，现有的对话、工具调用和流式展示仍然能够工作。

纯文本问答通常很容易接通。真正的复杂度会在下一步出现：模型同时调用两个工具，工具返回后继续回答，连接中途断开，或者用户希望恢复上一轮会话。此时，字段名称、消息顺序、事件生命周期都会影响结果。

协议转换的核心目标，是让客户端和模型服务对同一次交互形成一致的理解。本文以 OpenAI Chat Completions、OpenAI Responses 和 Anthropic Messages 的基础交互为例，讨论其中最容易出错的几个环节。具体模型支持哪些参数和工具，仍需单独确认。

## 三种 API 如何组织一次交互

可以先从三种 API 的基本组织单位看起。

| 维度 | Chat Completions | Responses | Messages |
| --- | --- | --- | --- |
| 输入主体 | `messages` 消息数组 | `input`，可包含多种 item | `messages`，消息内包含 content block |
| 指令入口 | `system` 或 `developer` 等消息角色，依模型而定 | `instructions` 及相应角色的输入消息 | 基础用法通过顶层 `system` 提供系统指令 |
| 普通输出 | `choices[].message` | `output[]` 中的 message item | assistant 消息的 text block |
| 函数调用 | assistant 消息的 `tool_calls[]` | 独立的 `function_call` item | assistant 消息中的 `tool_use` block |
| 函数结果 | `role: "tool"` 消息 | `function_call_output` item | user 消息中的 `tool_result` block |

Responses 把消息、函数调用等内容表达为不同类型的 item；Messages 则把不同内容组织在消息的 block 数组中。转换时，一条消息可能需要拆成多个对象，多个相邻对象也可能需要组合成一条消息。相关结构可参阅 [OpenAI 的 Responses 迁移说明](https://developers.openai.com/api/docs/guides/migrate-to-responses)和 [Anthropic Messages API](https://platform.claude.com/docs/en/api/messages/create)。

例如，一条 assistant 消息包含“我来查询一下”以及一个工具调用。转成 Responses 后，前者可以成为 message item，后者成为 function_call item。两部分仍然属于同一轮模型输出，它们的先后关系和关联信息都需要保留。

因此，开始编写字段映射之前，最好先回答三个问题：这段内容由谁产生，属于哪一轮交互，后续哪段内容会引用它。

## 请求转换需要保留内容的含义

### 指令和消息角色

系统指令、开发者指令、用户输入和工具结果具有不同的来源与作用。即使它们最终都包含文本，也不宜直接拼成一段字符串。

当目标协议只有较少的角色层级时，需要先确定角色映射策略。如果合并系统指令与开发者指令，应承认目标侧无法完整保留原来的优先级层次。用户输入和工具返回的数据则应继续留在各自的内容位置，避免在转换过程中被提升为系统指令。

同样，数组里的相邻元素也不一定对应不同轮次。可以合并属于同一语义单元的内容，但要保留工具调用与结果之间的因果关系。仅凭角色相同就合并，可能掩盖原本存在的交互边界。

### 文本和多模态内容

文本块之间通常比较容易转换，图片和文件还需要核对目标 API 的内容类型、媒体格式和模型能力。

以图片为例，某个协议可能使用带 `media_type` 的 base64 源，另一个协议可能接受 data URL。两者可以通过格式转换连接起来，但仍需检查 MIME 类型、大小和目标模型是否支持图片输入。

PDF 应使用目标端点支持的文件输入结构。例如，OpenAI Responses 使用 `input_file`，Chat Completions 的 PDF 输入使用 `file` 内容部分。把 PDF 的 base64 数据放进 `image_url`，不能保证目标服务会按文档处理它。[文件输入说明](https://developers.openai.com/api/docs/guides/file-inputs)列出了端点之间的差异。

### 参数名称相近也要核对范围

输出长度、采样参数和模型名都需要结合目标模型处理：

- 输出上限可能包含不同类别的 token，需要核对目标字段的定义。
- 某个 API 提供 `temperature`，不代表它的每个模型都接受相同配置。
- 模型名由供应商定义，跨供应商切换通常需要显式选择目标模型。

实用的做法是为每项能力标明处理方式：直接映射、经过验证的近似映射，或明确不支持。这样更容易解释转换后的行为。

认证配置也需要单独处理。跨供应商调用时，应使用目标服务认可的凭据；将认证头换一个名称，无法让原服务的密钥获得另一家服务的访问权限。

## 工具调用是一条需要闭合的链路

一次普通的客户端工具调用包含三个阶段：模型提出调用，应用执行工具，应用把结果交还给模型。协议转换需要同时覆盖工具定义、调用参数和结果回传。

### 同一次调用的三种表达

假设有一个自定义工具 `lookup_temperature`，参数为 `{"city":"杭州"}`。下面只展示与函数调用有关的结构片段，省略请求和响应的其他字段。

Chat Completions 将调用放在 assistant 消息中：

```json
{
  "role": "assistant",
  "content": null,
  "tool_calls": [
    {
      "id": "call_example_01",
      "type": "function",
      "function": {
        "name": "lookup_temperature",
        "arguments": "{\"city\":\"杭州\"}"
      }
    }
  ]
}
```

Messages 使用 `tool_use` 内容块：

```json
{
  "role": "assistant",
  "content": [
    {
      "type": "tool_use",
      "id": "call_example_01",
      "name": "lookup_temperature",
      "input": { "city": "杭州" }
    }
  ]
}
```

Responses 使用独立的 `function_call` item：

```json
{
  "type": "function_call",
  "id": "fc_example_01",
  "call_id": "call_example_01",
  "name": "lookup_temperature",
  "arguments": "{\"city\":\"杭州\"}"
}
```

这里有两个关键差异：`input` 是 JSON 对象，`arguments` 是承载 JSON 的字符串；Responses 的 item `id` 和关联工具结果的 `call_id` 各有用途。转换时需要解码或编码参数，并维护调用标识之间的对应关系。结构定义见 [OpenAI Function calling](https://developers.openai.com/api/docs/guides/function-calling)和 [Anthropic 工具调用处理](https://platform.claude.com/docs/en/agents-and-tools/tool-use/handle-tool-calls)。

### 工具结果要回到正确的调用

继续使用这个例子，假设测试工具返回字符串 `22°C`。结果回传时，对应关系如下：

| 协议 | 结果位置 | 关联字段 |
| --- | --- | --- |
| Chat Completions | `role: "tool"` 消息的 `content` | `tool_call_id: "call_example_01"` |
| Responses | `function_call_output` 的 `output` | `call_id: "call_example_01"` |
| Messages | user 消息中的 `tool_result.content` | `tool_use_id: "call_example_01"` |

同一轮出现多个调用时，不能只按返回顺序配对。例如两个城市的查询并发执行，第二个调用可能先完成。结果应始终通过调用 ID 找回对应请求。

工具结果还可能与新的用户文本出现在同一条消息中。Anthropic 要求结果消息紧跟对应的工具调用消息，并将 `tool_result` 块放在该消息普通文本之前。转换为 Chat Completions 时，可以先构造关联的 tool 消息，再表达剩余的用户文本，避免将工具结果当作用户自己的陈述。[工具结果的格式约束](https://platform.claude.com/docs/en/agents-and-tools/tool-use/handle-tool-calls#handling-results-from-client-tools)对此有明确说明。

### 工具定义和选择策略

普通函数工具常见的 Schema 对应关系是：

- Messages：`input_schema`。
- Chat Completions：`function.parameters`。
- Responses：function 工具的 `parameters`。

字段搬过去之后，还要检查 JSON Schema 特性的支持范围，以及严格校验模式的要求。工具名称如果需要转换，也应保留可逆映射，并处理重名，避免模型返回的名称无法找到原工具。

工具选择同样需要翻译语义。例如，在目标模型支持强制工具调用的前提下，OpenAI 的 `"required"` 可以对应 Anthropic 的 `{"type":"any"}`。具体枚举值和与推理模式的组合限制，需要按目标模型核对。[Anthropic 工具选择说明](https://platform.claude.com/docs/en/agents-and-tools/tool-use/define-tools)列出了这些约束。

供应商托管的搜索、代码执行等工具还包含服务端执行能力。转换它们的名称和参数，不会自动获得对应的执行环境。这类能力应单独适配。

## 流式转换需要维护事件生命周期

非流式响应可以在拿到完整 JSON 后统一转换。流式响应中，消息、文本块和工具参数会逐步形成，转换器需要记住已经发送了什么、哪些对象仍然打开、哪些对象已经结束。

### 先解析 SSE 再转换事件

网络读取边界与 SSE 事件边界彼此独立。一次读取可能只拿到半行，也可能拿到多个事件；UTF-8 字符也可能跨越读取边界。

合理的处理顺序是：

```text
网络字节
  → 增量解码
  → 按 SSE 规则识别行和事件
  → 解析事件中的 JSON
  → 更新协议状态
  → 输出目标协议事件
```

SSE 还允许多种换行形式、注释行以及同一事件中的多行 `data:`。解析器应按协议拼接数据，不能假定每次读取都得到完整的 `data: {...}`。[SSE 标准](https://html.spec.whatwg.org/multipage/server-sent-events.html#parsing-an-event-stream)定义了这些边界。

### 增量必须找到所属对象

以一次仅包含文本的正常回答为例，下面展示两种事件生命周期，省略事件体、可选事件和重复的增量：

```text
Anthropic Messages
  message_start
  content_block_start
  content_block_delta …
  content_block_stop
  message_delta
  message_stop

OpenAI Responses
  response.created
  response.in_progress
  response.output_item.added
  response.content_part.added
  response.output_text.delta …
  response.output_text.done
  response.content_part.done
  response.output_item.done
  response.completed
```

Messages 使用 block `index` 定位内容块，Responses 的相关事件使用 `item_id`、`output_index`、`content_index` 等字段定位对象，并通过 `sequence_number` 表达事件顺序。客户端需要先知道对象已经创建，才能将后续增量附加进去。两边的事件定义分别见 [Messages 流式文档](https://platform.claude.com/docs/en/build-with-claude/streaming)和 [Responses 事件参考](https://developers.openai.com/api/reference/resources/responses/streaming-events)。

因此，转换器通常需要维护三类信息：源对象到目标对象的对应关系，每个对象累计收到的内容，以及对象的生命周期状态。

例如，同一源消息的两个文本块被组合到一个目标 message item 中时，可以分别结束两个 content part，等整个 message 内容完整后再结束 item。如果 item 已经结束，后面又向它添加内容，客户端就可能产生与最终响应不同的结果。

### 工具参数要按调用分别拼接

流式工具参数通常以 JSON 字符串片段出现。一段增量可能只有半个字段名或一个尚未闭合的字符串，单独解析它会失败。

应为每个工具调用分别积累参数片段，等调用对应的参数流结束后，再进行完整 JSON 解析和 Schema 校验。参数不完整或不合法时，应报告错误，避免把它替换成空对象后继续执行工具。

这与普通文本增量的处理有明显区别：文本可以即时展示，工具执行需要等待完整且有效的输入。

### 文本结束后还可能有用量信息

Chat Completions 开启 `stream_options.include_usage` 后，会在 `[DONE]` 前额外发送一个用量 chunk，其 `choices` 为空数组。收到这个 chunk 时，不能继续假设存在 `choices[0]`。[Chat Completions 参考](https://developers.openai.com/api/reference/resources/chat/subresources/completions/methods/create)同时说明，流中断时可能收不到最后的用量统计。

转换器应分别跟踪“生成内容已结束”和“整个响应已收尾”。这样既能及时结束文本块，也能在获得最终统计后完成目标响应。

## 完成失败和中断需要分别表达

模型正常结束回答、达到输出上限、请求执行工具，以及网络连接断开，表示不同的状态。

Responses 提供 `response.completed`、`response.incomplete` 和 `response.failed` 等事件。选择目标状态时，需要依据上游实际发生的事情，并满足对应的事件 Schema。[Responses 事件参考](https://developers.openai.com/api/reference/resources/responses/streaming-events)定义了各类终态。

工具调用可以是这一轮模型响应的正常结果；应用仍需要执行工具并发起后续请求。达到输出上限则需要保留截断含义，让调用方决定是否继续生成。

网络 EOF 只说明数据流已经结束。若此前没有收到有效的完成信号，就无法据此证明模型完成了回答。此时应按客户端协议报告流错误或未完成状态，并保留已经接收到的内容。连接已经断开、无法再发送事件时，由调用方记录这一事实。

错误转换也应覆盖两个阶段：响应尚未开始时，可以通过 HTTP 状态和错误体表达；流式响应已经开始后，需要使用流内错误机制。两种情况下，都应避免同时给出失败和成功完成的相互矛盾信号。

## 推理与会话状态存在转换边界

### 推理强度只能按模型适配

OpenAI 的 reasoning effort 表示推理投入档位，具体可选值依模型而异。Anthropic 的推理控制也随模型与模式变化；支持手动预算的模式可以使用 token budget，自适应模式有自己的配置方式。[OpenAI 推理说明](https://developers.openai.com/api/docs/guides/reasoning)和 [Anthropic 推理说明](https://platform.claude.com/docs/en/build-with-claude/extended-thinking)分别描述了这些控制方式。

档位与 token 数量之间没有跨供应商统一的换算公式。若要建立近似映射，应将它作为明确的适配策略，并检查输出预算、最小值和工具调用配置是否允许该组合。

推理的可展示文本、摘要、签名和加密状态也各有用途。带签名或加密的供应商状态应按来源协议保存和回传，不能仅因字段类型相同，就视为另一家供应商可接受的状态。

### 会话引用依赖可恢复的历史

Responses 可以通过 `previous_response_id` 引用前一次响应，也可以由应用维护并回传所需的历史 item。[会话迁移说明](https://developers.openai.com/api/docs/guides/migrate-to-responses)介绍了这些方式。

如果转换层只收到一个响应 ID，且没有保存相关历史，也无法向原服务取回历史，就无法重建完整上下文。转换后的请求即使能被目标服务接受，也可能已经失去上一轮的工具调用和回答。

因此，应在能力约定中说明采用哪一种方式：由调用方提供完整历史，或由服务保存并恢复上下文。无法满足的会话恢复请求，应明确返回不支持。

即使保存了历史，跨模型恢复也要检查其中的多模态数据、工具类型和供应商状态是否能被目标端接受。

## 用量统计要先对齐定义

输入 token、输出 token、缓存命中 token 和推理 token 可能采用不同的统计口径。相似的字段名不保证数值可以直接替换。

例如，Anthropic 的缓存统计将输入分为普通输入、缓存读取和缓存创建。要表达完整输入量，需要结合这几个字段计算：[Prompt caching 用量说明](https://platform.claude.com/docs/en/build-with-claude/prompt-caching)。

```text
完整输入量 = input_tokens
           + cache_read_input_tokens
           + cache_creation_input_tokens
```

如果统一接口还需要缓存命中量，就应另外保留缓存读取明细，避免在合计时重复相加。流式统计还要区分累计值与本次增量，累计值应更新已有统计。

上游没有报告某个明细时，应保留“未知”或“不可用”的含义，并遵守目标 Schema 的可选字段约束。把未知值写成零，会让调用方误以为已经确认没有这部分消耗。

这些数值来自实际执行请求的模型，其 tokenizer 和计费规则仍属于原供应商。转换后的统计字段适合统一展示，账单核对还需要保留供应商的原始口径。

## 如何验证一次转换是否正确

验证可以围绕客户端最终能够观察到的行为展开。

| 场景 | 应检查的结果 |
| --- | --- |
| 多轮纯文本 | 角色、文本和轮次顺序保持一致 |
| 单个工具调用 | 工具定义、参数、调用 ID 和结果形成完整链路 |
| 多个并行工具 | 参数与结果不会因交错到达而串到其他调用 |
| 文本与工具混合 | 两种内容均保留，并满足目标协议的排列约束 |
| 随机拆分网络数据 | SSE 解析结果不受读取边界影响 |
| 输出达到上限 | 截断状态能被客户端识别 |
| 流式异常与断连 | 已有内容可保留，结束状态不会被误报为成功 |
| 上下文恢复 | 历史完整，调用与结果仍能正确关联 |
| 不支持的内容或工具 | 按公开约定拒绝或降级，调用方能够感知差异 |

还有一个很有价值的检查方法：对同一份固定的逻辑输出，分别生成完整响应和流式事件，再让目标客户端聚合事件，比较两者的文本、工具参数、调用 ID、结束状态和用量。这能发现“屏幕上看到了文字，但保存下来的回答不完整”一类问题。

这种验证应使用固定数据或可控的事件回放。分别请求模型两次，结果可能因生成随机性而不同，难以据此判断转换是否正确。

完成这些检查后，还需要让实际客户端走完一次“提问、调用工具、接收结果、继续回答”的完整过程。协议字段、流式状态与会话历史在这个过程中都能被检验，转换能力的边界也会更清楚。
