# 51. N 皇后

分类：回溯 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/n-queens/)

**题意：** 在 n×n 棋盘上放置 n 个皇后，使任意两者不同行、不同列、不同对角线，返回全部棋盘方案。

**示例：** `n = 4` → `[[".Q..","...Q","Q...","..Q."],["..Q.","Q...","...Q",".Q.."]]`。

## 方法、入参与返回值

### `solveNQueens`

```go
func solveNQueens(n int) [][]string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `n` | `int` | 棋盘边长及皇后数量 |

返回值类型：`[][]string`。所有合法棋盘方案，用 Q 表示皇后、点号表示空位；无解时返回空列表。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0051_n_queens -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
