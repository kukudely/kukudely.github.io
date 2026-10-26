# 240. 搜索二维矩阵 II

分类：矩阵 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/search-a-2d-matrix-ii/)

**题意：** 给定每行从左到右递增、每列从上到下递增的矩阵，判断目标值是否存在。

**示例：** matrix = [[1,4,7],[2,5,8],[3,6,9]], target = 6 → true。

## 方法、入参与返回值

### `searchMatrix`

```go
func searchMatrix(matrix [][]int, target int) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `matrix` | `[][]int` | 每行、每列分别递增的整数矩阵 |
| `target` | `int` | 待查找数值 |

返回值类型：`bool`。目标值存在时为 true，否则为 false。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0240_search_a_2d_matrix_ii -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
