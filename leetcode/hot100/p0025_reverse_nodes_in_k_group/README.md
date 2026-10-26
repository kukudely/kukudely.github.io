# 25. K 个一组翻转链表

分类：链表 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/reverse-nodes-in-k-group/)

**题意：** 每 k 个节点为一组反转链表，不足 k 个的尾部保持原顺序；只能调整节点连接。

**示例：** head = [1,2,3,4,5], k = 2 → [2,1,4,3,5]。

## 方法、入参与返回值

### `reverseKGroup`

```go
func reverseKGroup(head *ListNode, k int) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点 |
| `k` | `int` | 每组节点数，k ≥ 1 |

返回值类型：`*ListNode`。每 k 个节点一组反转后的头节点；末尾不足 k 个节点保持原顺序。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0025_reverse_nodes_in_k_group -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
