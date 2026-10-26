# 208. 实现 Trie (前缀树)

分类：图论 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/implement-trie-prefix-tree/)

**题意：** 实现插入单词、查询完整单词和查询前缀。题目限定小写英文字母。下面两份代码分别提交，类型及 Constructor 同名。

**示例：** `Insert("apple")` 后，`Search("apple") = true`、`Search("app") = false`、`StartsWith("app") = true`；再插入 `"app"` 后可查询到完整单词。

## 方法、入参与返回值

### `Constructor`

```go
func Constructor() Trie
```

入参：无。

返回值类型：`Trie`。新建空 Trie。

### `Insert`

```go
func (t *Trie) Insert(word string)
```

接收者：`*Trie`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `word` | `string` | 待插入的完整单词 |

返回值类型：`无`。无返回值；记录单词，允许重复插入。

### `Search`

```go
func (t *Trie) Search(word string) bool
```

接收者：`*Trie`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `word` | `string` | 待查找的完整单词 |

返回值类型：`bool`。只有完整单词已插入时才返回 true。

### `StartsWith`

```go
func (t *Trie) StartsWith(prefix string) bool
```

接收者：`*Trie`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `prefix` | `string` | 待检查的前缀 |

返回值类型：`bool`。存在具有该前缀的已插入单词时返回 true；空前缀按测试约定为 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0208_implement_trie_prefix_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
