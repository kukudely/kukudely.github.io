# 238. 除了自身以外数组的乘积

分类：普通数组 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/product-of-array-except-self/)

**题意：** 返回每个位置除自身以外所有元素的乘积，题目禁止除法，并要求 O(n) 时间。

**示例：** nums = [1,2,3,4] → [24,12,8,6]。

## 方法、入参与返回值

### `productExceptSelf`

```go
func productExceptSelf(nums []int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |

返回值类型：`[]int`。answer[i] 为除 nums[i] 外其余元素的乘积；无需使用除法。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0238_product_of_array_except_self -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
