# 22. 括号生成

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/generate-parentheses/)

**题意：** 生成 n 对括号能组成的全部合法括号字符串。

**示例：** `n = 3` → `["((()))","(()())","(())()","()(())","()()()"]`。

## 方法、入参与返回值

### `generateParenthesis`

```go
func generateParenthesis(n int) []string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `n` | `int` | 括号对数 |

返回值类型：`[]string`。由 n 对括号构成的全部合法字符串，结果顺序不限。

## 本地测试约定

额外边界：n=0 时返回包含一个空字符串的列表，即 []string{""}。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0022_generate_parentheses -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
