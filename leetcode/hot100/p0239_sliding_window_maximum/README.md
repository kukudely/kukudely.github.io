# 239. 滑动窗口最大值

分类：子串 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/sliding-window-maximum/)

**题意：** 窗口长度为 k，从左到右每次移动一格，输出每个窗口的最大值。

**示例：** nums = [1,3,-1,-3,5,3,6,7], k = 3 → [3,3,5,5,6,7]。

## 方法、入参与返回值

### `maxSlidingWindow`

```go
func maxSlidingWindow(nums []int, k int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |
| `k` | `int` | 窗口大小，1 ≤ k ≤ len(nums) |

返回值类型：`[]int`。每个连续窗口的最大值，按窗口从左到右排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0239_sliding_window_maximum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
