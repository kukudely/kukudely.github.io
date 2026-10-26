# 84. 柱状图中最大的矩形

分类：栈 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/largest-rectangle-in-histogram/)

**题意**：柱状图每根柱子的宽度为 1，高度非负，求由连续柱子形成的最大矩形面积。

**示例**：heights = [2,1,5,6,2,3] → 10。

## 方法、入参与返回值

### `largestRectangleArea`

```go
func largestRectangleArea(heights []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `heights` | `[]int` | 柱状图高度列表，每根柱宽为 1 |

返回值类型：`int`。连续柱子可以形成的最大矩形面积。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0084_largest_rectangle_in_histogram -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
