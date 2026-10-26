# 560. 和为 K 的子数组

分类：子串 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/subarray-sum-equals-k/)

**题意：** 统计元素和恰好为 k 的连续非空子数组数量。数组允许负数和零。

**示例：** nums = [1,1,1], k = 2 → 2。

## 方法、入参与返回值

### `subarraySum`

```go
func subarraySum(nums []int, k int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组，可以含负数和零 |
| `k` | `int` | 目标子数组和 |

返回值类型：`int`。和等于 k 的非空连续子数组数量。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0560_subarray_sum_equals_k -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
