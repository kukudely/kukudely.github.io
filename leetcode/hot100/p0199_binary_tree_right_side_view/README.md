# 199. 二叉树的右视图

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/binary-tree-right-side-view/)

**题意：** 从树的右侧观察，按深度输出每层最右侧可见节点的值。

**示例：** `root = [1,2,3,null,5,null,4]` → `[1,3,4]`。

## 方法、入参与返回值

### `rightSideView`

```go
func rightSideView(root *TreeNode) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`[]int`。从右侧观察时每一层能看到的节点值，自上而下排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0199_binary_tree_right_side_view -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
