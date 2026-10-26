# 200. 岛屿数量

分类：图论 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/number-of-islands/)

**题意：** 在字符网格中统计由上下左右相连的 '1' 组成的连通块，'0' 表示水。

**示例：** `grid = ["11000","11000","00100","00011"]` → `3`。

## 方法、入参与返回值

### `numIslands`

```go
func numIslands(grid [][]byte) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `grid` | `[][]byte` | 由字节字符 0 和 1 组成的二维网格 |

返回值类型：`int`。四方向相邻的陆地连通块数量。注意输入元素为字符字节。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0200_number_of_islands -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
