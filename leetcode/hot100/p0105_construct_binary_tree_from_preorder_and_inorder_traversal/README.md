# 105. 从前序与中序遍历序列构造二叉树

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)

**题意：** 给定同一棵树的前序和中序遍历，重建二叉树。题目保证节点值互不相同且遍历有效。

**示例：** `preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]` → `[3,9,20,null,null,15,7]`。

## 方法、入参与返回值

### `buildTree`

```go
func buildTree(preorder []int, inorder []int) *TreeNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `preorder` | `[]int` | 节点值互异的前序遍历 |
| `inorder` | `[]int` | 同一棵树的中序遍历 |

返回值类型：`*TreeNode`。由两种遍历重建的二叉树根节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0105_construct_binary_tree_from_preorder_and_inorder_traversal -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
