# 994. 腐烂的橘子

分类：图论 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/rotting-oranges/)

**题意：** 0 为空格、1 为新鲜橘子、2 为腐烂橘子。每分钟腐烂向四邻传播，求全部新鲜橘子腐烂的最少分钟数；无法完成时返回 -1。两种实现均修改网格。

**示例：** `grid = [[2,1,1],[1,1,0],[0,1,1]]` → `4`。

## 方法、入参与返回值

### `orangesRotting`

```go
func orangesRotting(grid [][]int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `grid` | `[][]int` | 0 为空格、1 为新鲜橘子、2 为腐烂橘子的二维网格 |

返回值类型：`int`。使全部橘子腐烂的最少分钟数；无法全部腐烂时返回 -1，无新鲜橘子时返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0994_rotting_oranges -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
