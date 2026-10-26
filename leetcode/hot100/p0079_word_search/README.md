# 79. 单词搜索

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/word-search/)

**题意：** 判断矩形字母网格中是否存在一条上下左右相邻的路径拼出 word，同一格在一条路径中只能使用一次。

**示例：** `board = ["ABCE","SFCS","ADEE"], word = "ABCCED"` → `true`；`word = "ABCB"` → `false`。

## 方法、入参与返回值

### `exist`

```go
func exist(board [][]byte, word string) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `board` | `[][]byte` | 字符网格 |
| `word` | `string` | 待搜索单词 |

返回值类型：`bool`。存在连续上下左右路径拼出 word 时返回 true；一条路径不得重复使用格子。

## 本地测试约定

测试会检查搜索后网格恢复；在网格上临时标记访问状态时，请在回溯时撤销。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0079_word_search -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
