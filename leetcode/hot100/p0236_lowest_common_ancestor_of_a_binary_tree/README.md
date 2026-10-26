# 236. 二叉树的最近公共祖先

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/)

**题意：** 返回给定两个节点 p、q 的最近公共祖先。节点可以是自己的祖先；题目保证两节点都在树内。

**示例：** `root = [3,5,1,6,2,0,8,null,null,7,4], p = 5, q = 1` → 节点 `3`。

## 方法、入参与返回值

### `lowestCommonAncestor`

```go
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |
| `p` | `*TreeNode` | 树中存在的第一个目标节点指针 |
| `q` | `*TreeNode` | 树中存在的第二个目标节点指针 |

返回值类型：`*TreeNode`。p、q 的最近公共祖先原节点指针；节点也可以是自身的祖先。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0236_lowest_common_ancestor_of_a_binary_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
