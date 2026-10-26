# 543. 二叉树的直径

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/diameter-of-binary-tree/)

**题意：** 求任意两节点间最长路径的边数，路径可以经过树中任意节点。

**示例：** `root = [1,2,3,4,5]` → `3`。

## 方法、入参与返回值

### `diameterOfBinaryTree`

```go
func diameterOfBinaryTree(root *TreeNode) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`int`。任意两个节点间最长路径的边数；空树返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0543_diameter_of_binary_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
