# 394. 字符串解码

分类：栈 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/decode-string/)

**题意**：解码形如 k[encoded_string] 的字符串；方括号可以嵌套，k 可以是多位正整数。

**示例**：s = "3[a2[c]]" → "accaccacc"。

## 方法、入参与返回值

### `decodeString`

```go
func decodeString(s string) string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 格式合法的编码字符串，形式如 3[a2[c]] |

返回值类型：`string`。将 k[片段] 递归展开 k 次后的字符串。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0394_decode_string -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
