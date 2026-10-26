# 139. 单词拆分

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/word-break/)

**题意**：判断字符串是否能被拆成字典中的若干单词。字典单词可以重复使用。

**示例**：s = "leetcode", wordDict = ["leet","code"] → true。

## 方法、入参与返回值

### `wordBreak`

```go
func wordBreak(s string, wordDict []string) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 待拆分字符串 |
| `wordDict` | `[]string` | 由非空小写字符串组成的字典，单词可重复使用 |

返回值类型：`bool`。s 能完整拆分为字典中的若干单词时返回 true。

## 本地测试约定

额外边界：空字符串无需选择任何单词即可完成拆分，返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0139_word_break -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
