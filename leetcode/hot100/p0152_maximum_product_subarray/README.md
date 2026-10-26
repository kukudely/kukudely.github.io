# 152. 乘积最大子数组

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/maximum-product-subarray/)

**题意**：求非空连续子数组能够取得的最大乘积。数组可能包含正数、负数和零。

**示例**：nums = [2,3,-2,4] → 6；nums = [-2,0,-1] → 0。

## 方法、入参与返回值

### `maxProduct`

```go
func maxProduct(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非空整数数组 |

返回值类型：`int`。非空连续子数组乘积的最大值。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0152_maximum_product_subarray -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
