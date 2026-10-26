# 33. 搜索旋转排序数组

分类：二分查找 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/search-in-rotated-sorted-array/)

**题意**：一个严格升序数组经过旋转，元素互不相同。返回 target 所在下标；不存在时返回 -1。

**示例**：nums = [4,5,6,7,0,1,2], target = 0 → 4。

## 方法、入参与返回值

### `search`

```go
func search(nums []int, target int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 由严格递增数组旋转得到的数组，元素互异 |
| `target` | `int` | 目标值 |

返回值类型：`int`。目标值所在下标；不存在时返回 -1。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0033_search_in_rotated_sorted_array -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
