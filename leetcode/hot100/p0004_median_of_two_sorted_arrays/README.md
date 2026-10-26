# 4. 寻找两个正序数组的中位数

分类：二分查找 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/median-of-two-sorted-arrays/)

**题意**：两个非递减数组总长度至少为 1，求合并后有序序列的中位数。

**示例**：nums1 = [1,2], nums2 = [3,4] → 2.5。

## 方法、入参与返回值

### `findMedianSortedArrays`

```go
func findMedianSortedArrays(nums1 []int, nums2 []int) float64
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums1` | `[]int` | 第一个非递减整数数组 |
| `nums2` | `[]int` | 第二个非递减整数数组，两数组总长度大于 0 |

返回值类型：`float64`。两个数组合并后序列的中位数；偶数长度时返回中间两数的平均值。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0004_median_of_two_sorted_arrays -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
