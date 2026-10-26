# 1143. 最长公共子序列

分类：多维动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/longest-common-subsequence/)

**题意**：返回两个字符串的最长公共子序列长度。子序列保留相对顺序，可以删除中间字符。

**示例**：text1 = "abcde", text2 = "ace" → 3。

## 方法、入参与返回值

### `longestCommonSubsequence`

```go
func longestCommonSubsequence(text1 string, text2 string) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `text1` | `string` | 第一个字符串 |
| `text2` | `string` | 第二个字符串 |

返回值类型：`int`。两个字符串最长公共子序列的长度；子序列无需连续。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p1143_longest_common_subsequence -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
