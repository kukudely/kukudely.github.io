# 5. 最长回文子串

分类：多维动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/longest-palindromic-substring/)

**题意**：返回字符串中最长的连续回文子串。有多个最长答案时，返回其中任意一个。

**示例**：s = "babad" → "bab" 或 "aba"；s = "cbbd" → "bb"。

## 方法、入参与返回值

### `longestPalindrome`

```go
func longestPalindrome(s string) string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 待检查字符串 |

返回值类型：`string`。最长回文连续子串；存在多个最长答案时任意一个均可。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0005_longest_palindromic_substring -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
