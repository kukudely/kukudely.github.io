# 287. 寻找重复数

分类：技巧 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/find-the-duplicate-number/)

**题意**：长度为 n+1 的数组中，数值均在 1..n，只有一种数值出现多次，找出它。要求不修改数组，并且额外空间为 O(1)。

**示例**：nums = [1,3,4,2,2] → 2。

## 方法、入参与返回值

### `findDuplicate`

```go
func findDuplicate(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 长度为 n+1、元素范围为 [1,n] 的整数数组，仅一个不同数值重复 |

返回值类型：`int`。重复的数值；该值可能出现两次以上。

## 本地测试约定

原题进阶要求保持输入数组不变且仅使用常数额外空间；基础功能测试无法证明复杂度。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0287_find_the_duplicate_number -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
