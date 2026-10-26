# 226. 翻转二叉树

分类：二叉树 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/invert-binary-tree/)

**题意：** 交换每个节点的左右子树，返回翻转后的根节点。以下两种写法都会原地修改树。

**示例：** `root = [4,2,7,1,3,6,9]` → `[4,7,2,9,6,3,1]`。

## 方法、入参与返回值

### `invertTree`

```go
func invertTree(root *TreeNode) *TreeNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 需要翻转的二叉树根节点 |

返回值类型：`*TreeNode`。翻转左右子树后的根节点，复用原树节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0226_invert_binary_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
