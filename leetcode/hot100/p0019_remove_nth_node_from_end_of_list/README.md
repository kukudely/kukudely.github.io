# 19. 删除链表的倒数第 N 个结点

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/remove-nth-node-from-end-of-list/)

**题意：** 删除链表倒数第 n 个节点，返回头节点；题目保证 1≤n≤链表长度。

**示例：** head = [1,2,3,4,5], n = 2 → [1,2,3,5]。

## 方法、入参与返回值

### `removeNthFromEnd`

```go
func removeNthFromEnd(head *ListNode, n int) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点 |
| `n` | `int` | 需要删除节点的倒数序号，从 1 开始，保证有效 |

返回值类型：`*ListNode`。删除倒数第 n 个节点后的头节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0019_remove_nth_node_from_end_of_list -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
