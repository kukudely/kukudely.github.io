# 72. 编辑距离

分类：多维动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/edit-distance/)

**题意**：将 word1 变成 word2，可插入、删除或替换一个字符，求最少操作数。

**示例**：word1 = "horse", word2 = "ros" → 3。

## 方法、入参与返回值

### `minDistance`

```go
func minDistance(word1 string, word2 string) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `word1` | `string` | 原始单词 |
| `word2` | `string` | 目标单词 |

返回值类型：`int`。将 word1 变为 word2 的最少插入、删除或替换次数，每次操作代价为 1。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0072_edit_distance -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
