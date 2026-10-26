# 230. 二叉搜索树中第 K 小的元素

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/kth-smallest-element-in-a-bst/)

**题意：** 返回 BST 中第 k 小的节点值，k 从 1 开始。题目保证 k 有效。

**示例：** `root = [3,1,4,null,2], k = 1` → `1`。

## 方法、入参与返回值

### `kthSmallest`

```go
func kthSmallest(root *TreeNode, k int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉搜索树根节点 |
| `k` | `int` | 从 1 开始的序号，保证有效 |

返回值类型：`int`。中序遍历中的第 k 小节点值。

## 本地测试约定

测试会检查查询后原树保持完整；提前返回前须处理临时建立的连接。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0230_kth_smallest_element_in_a_bst -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
