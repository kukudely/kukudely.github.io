# 64. 最小路径和

分类：多维动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/minimum-path-sum/)

**题意**：非负整数网格中，每次只能向右或向下，求从左上角到右下角的最小路径和。

**示例**：grid = [[1,3,1],[1,5,1],[4,2,1]] → 7。

## 方法、入参与返回值

### `minPathSum`

```go
func minPathSum(grid [][]int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `grid` | `[][]int` | 非负整数网格 |

返回值类型：`int`。从左上角仅向右或向下到右下角的最小路径和，包含起终点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0064_minimum_path_sum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
