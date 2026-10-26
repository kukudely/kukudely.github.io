# 53. 最大子数组和

分类：普通数组 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/maximum-subarray/)

**题意：** 在非空整数数组中选择一个连续非空子数组，使元素和最大。

**示例：** nums = [-2,1,-3,4,-1,2,1,-5,4] → 6，对应 [4,-1,2,1]。

## 方法、入参与返回值

### `maxSubArray`

```go
func maxSubArray(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非空整数数组 |

返回值类型：`int`。非空连续子数组能够取得的最大元素和。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0053_maximum_subarray -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
