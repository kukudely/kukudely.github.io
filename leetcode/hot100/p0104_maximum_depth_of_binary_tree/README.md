# 104. 二叉树的最大深度

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/maximum-depth-of-binary-tree/)

**题意：** 返回从根节点到最深叶子的节点数；空树深度为 0。

**示例：** `root = [3,9,20,null,null,15,7]` → `3`。

## 方法、入参与返回值

### `maxDepth`

```go
func maxDepth(root *TreeNode) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`int`。最长根到叶路径的节点数；空树返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0104_maximum_depth_of_binary_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
