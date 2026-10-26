# 98. 验证二叉搜索树

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/validate-binary-search-tree/)

**题意：** 判断每个节点是否都满足左子树所有值小于它、右子树所有值大于它；重复值不允许。

**示例：** `root = [2,1,3]` → `true`；`[5,1,4,null,null,3,6]` → `false`。

## 方法、入参与返回值

### `isValidBST`

```go
func isValidBST(root *TreeNode) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |

返回值类型：`bool`。所有节点满足严格二叉搜索树关系时返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0098_validate_binary_search_tree -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
