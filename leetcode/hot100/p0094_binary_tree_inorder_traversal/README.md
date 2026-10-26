# 94. 二叉树的中序遍历

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/binary-tree-inorder-traversal/)

**题意：** 按左子树、根节点、右子树的顺序返回节点值。

**示例：** `root = [1,null,2,3]` → `[1,3,2]`。

## 方法、入参与返回值

### `inorderTraversal`

```go
func inorderTraversal(root *TreeNode) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点，可以为 nil |

返回值类型：`[]int`。左子树、根、右子树顺序的节点值列表。

## 本地测试约定

测试会检查遍历结束后原树结构保持完整；使用临时线索时须拆除。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0094_binary_tree_inorder_traversal -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
