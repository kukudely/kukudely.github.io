# 74. 搜索二维矩阵

分类：二分查找 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/search-a-2d-matrix/)

**题意**：矩阵每行非递减，且下一行首元素严格大于上一行末元素。判断 target 是否存在。

**示例**：matrix = [[1,3,5,7],[10,11,16,20],[23,30,34,60]], target = 3 → true。

## 方法、入参与返回值

### `searchMatrix`

```go
func searchMatrix(matrix [][]int, target int) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `matrix` | `[][]int` | 每行非递减且下一行首元素大于上一行末元素的整数矩阵 |
| `target` | `int` | 目标值 |

返回值类型：`bool`。矩阵中存在 target 时返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0074_search_a_2d_matrix -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
