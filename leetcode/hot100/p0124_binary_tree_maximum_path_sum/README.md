# 124. 二叉树中的最大路径和

分类：二叉树 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)

**题意：** 求一条非空简单路径的最大节点值之和，路径可以在任意节点起止。

**示例：** `root = [-10,9,20,null,null,15,7]` → `42`，路径为 `15→20→7`。

## 方法、入参与返回值

### `maxPathSum`

```go
func maxPathSum(root *TreeNode) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点，原题保证非空 |

返回值类型：`int`。任意非空简单路径的最大节点值之和。

## 本地测试约定

为便于本地边界练习，额外约定空树返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0124_binary_tree_maximum_path_sum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
