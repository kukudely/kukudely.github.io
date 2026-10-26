# 136. 只出现一次的数字

分类：技巧 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/single-number/)

**题意**：数组中只有一个数出现一次，其他数都出现两次，找出这个单独的数。要求线性时间和常数额外空间。

**示例**：nums = [4,1,2,1,2] → 4。

## 方法、入参与返回值

### `singleNumber`

```go
func singleNumber(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 除一个值出现一次外，其余值均出现两次的整数数组 |

返回值类型：`int`。唯一只出现一次的元素值。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0136_single_number -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
