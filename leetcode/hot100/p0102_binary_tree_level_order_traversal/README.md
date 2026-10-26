# 102. 二叉树的层序遍历

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/binary-tree-level-order-traversal/)

**题意：** 从上到下逐层输出二叉树节点，每层按从左到右排列。

**示例：** `root = [3,9,20,null,null,15,7]` → `[[3],[9,20],[15,7]]`。

## 方法、入参与返回值

### `levelOrder`

```go
func levelOrder(root *TreeNode) [][]int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`[][]int`。每层一组、自上而下的节点值；每层从左到右排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0102_binary_tree_level_order_traversal -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
