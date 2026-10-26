# 215. 数组中的第K个最大元素

分类：堆 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/kth-largest-element-in-an-array/)

**题意**：返回数组按降序排列后的第 k 个元素；重复元素按出现次数计入排名。

**示例**：nums = [3,2,1,5,6,4], k = 2 → 5。

## 方法、入参与返回值

### `findKthLargest`

```go
func findKthLargest(nums []int, k int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |
| `k` | `int` | 按非递增排序后的序号，从 1 开始 |

返回值类型：`int`。第 k 大元素的值，重复元素分别计数。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0215_kth_largest_element_in_an_array -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
