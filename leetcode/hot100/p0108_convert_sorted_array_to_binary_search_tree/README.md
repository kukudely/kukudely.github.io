# 108. 将有序数组转换为二叉搜索树

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/convert-sorted-array-to-binary-search-tree/)

**题意：** 把严格递增数组构造成高度平衡的二叉搜索树，左右子树高度差最多为 1。答案可能有多种。

**示例：** `nums = [-10,-3,0,5,9]` → 一种答案为 `[0,-3,9,-10,null,5]`。

## 方法、入参与返回值

### `sortedArrayToBST`

```go
func sortedArrayToBST(nums []int) *TreeNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 严格递增整数数组 |

返回值类型：`*TreeNode`。包含全部输入值的高度平衡二叉搜索树根节点；可返回任意满足要求的形状。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0108_convert_sorted_array_to_binary_search_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
