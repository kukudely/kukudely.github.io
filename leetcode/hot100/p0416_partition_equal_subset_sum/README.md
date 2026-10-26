# 416. 分割等和子集

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/partition-equal-subset-sum/)

**题意**：把正整数数组分成两个子集，使两边元素和相等，判断是否可行。

**示例**：nums = [1,5,11,5] → true；[1,2,3,5] → false。

## 方法、入参与返回值

### `canPartition`

```go
func canPartition(nums []int) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 正整数数组 |

返回值类型：`bool`。能将全部元素分成和相等的两个子集时返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0416_partition_equal_subset_sum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
