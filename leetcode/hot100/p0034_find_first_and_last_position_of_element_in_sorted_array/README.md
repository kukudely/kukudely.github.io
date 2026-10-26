# 34. 在排序数组中查找元素的第一个和最后一个位置

分类：二分查找 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/find-first-and-last-position-of-element-in-sorted-array/)

**题意**：在非递减数组中找 target 第一次和最后一次出现的下标。不存在时返回 [-1,-1]。

**示例**：nums = [5,7,7,8,8,10], target = 8 → [3,4]。

## 方法、入参与返回值

### `searchRange`

```go
func searchRange(nums []int, target int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非递减整数数组 |
| `target` | `int` | 目标值 |

返回值类型：`[]int`。目标值首次及最后出现的下标；不存在时返回 [-1,-1]。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0034_find_first_and_last_position_of_element_in_sorted_array -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
