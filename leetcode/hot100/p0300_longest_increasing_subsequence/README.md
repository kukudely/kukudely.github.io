# 300. 最长递增子序列

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/longest-increasing-subsequence/)

**题意**：求数组中最长严格递增子序列的长度。子序列可以跳过元素，但下标顺序必须保持。

**示例**：nums = [10,9,2,5,3,7,101,18] → 4。

## 方法、入参与返回值

### `lengthOfLIS`

```go
func lengthOfLIS(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |

返回值类型：`int`。严格递增子序列的最大长度；子序列无需连续。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0300_longest_increasing_subsequence -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
