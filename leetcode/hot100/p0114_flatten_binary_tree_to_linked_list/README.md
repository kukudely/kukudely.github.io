# 114. 二叉树展开为链表

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/flatten-binary-tree-to-linked-list/)

**题意：** 原地把二叉树按前序顺序展开，所有 Left 指针置空，Right 指向后继。

**示例：** `root = [1,2,5,3,4,null,6]` → 右链 `1→2→3→4→5→6`。

## 方法、入参与返回值

### `flatten`

```go
func flatten(root *TreeNode)
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 需要原地展开的二叉树根节点 |

返回值类型：`无`。无返回值；按前序顺序用 Right 串联原节点，全部 Left 置为 nil。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0114_flatten_binary_tree_to_linked_list -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
