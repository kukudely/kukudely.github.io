# 20. 有效的括号

分类：栈 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/valid-parentheses/)

**题意**：字符串只包含 ()[]{}，判断每个左括号能否按正确嵌套顺序与同类右括号配对。

**示例**：s = "([])" → true；s = "([)]" → false。

## 方法、入参与返回值

### `isValid`

```go
func isValid(s string) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 仅包含六种圆括号、方括号、花括号字符的字符串 |

返回值类型：`bool`。括号类型和闭合顺序均正确时返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0020_valid_parentheses -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
