# 101. 对称二叉树

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/symmetric-tree/)

**题意：** 判断一棵二叉树是否关于根节点左右镜像对称。

**示例：** `root = [1,2,2,3,4,4,3]` → `true`。

## 方法、入参与返回值

### `isSymmetric`

```go
func isSymmetric(root *TreeNode) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`bool`。树关于根节点镜像对称时返回 true；空树视为对称。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0101_symmetric_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
