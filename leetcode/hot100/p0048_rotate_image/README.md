# 48. 旋转图像

分类：矩阵 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/rotate-image/)

**题意：** 将 n×n 方阵顺时针旋转90度，题目要求原地完成。

**示例：** matrix = [[1,2,3],[4,5,6],[7,8,9]] → [[7,4,1],[8,5,2],[9,6,3]]。

## 方法、入参与返回值

### `rotate`

```go
func rotate(matrix [][]int)
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `matrix` | `[][]int` | 需要原地修改的 n×n 整数方阵 |

返回值类型：`无`。无返回值；将 matrix 原地顺时针旋转 90 度。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0048_rotate_image -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
