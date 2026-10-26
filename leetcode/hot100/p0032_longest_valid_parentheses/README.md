# 32. 最长有效括号

分类：动态规划 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/longest-valid-parentheses/)

**题意**：给定只包含左右括号的字符串，求最长连续有效括号子串的长度。

**示例**：s = ")()())" → 4；s = "(()" → 2。

## 方法、入参与返回值

### `longestValidParentheses`

```go
func longestValidParentheses(s string) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 仅含左、右圆括号的字符串 |

返回值类型：`int`。最长连续合法括号子串的长度。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0032_longest_valid_parentheses -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
