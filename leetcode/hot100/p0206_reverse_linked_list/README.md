# 206. 反转链表

分类：链表 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/reverse-linked-list/)

**题意：** 反转单链表的 Next 指向，返回新的头节点。

**示例：** head = [1,2,3,4,5] → [5,4,3,2,1]。

## 方法、入参与返回值

### `reverseList`

```go
func reverseList(head *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点，可以为 nil |

返回值类型：`*ListNode`。反转后链表的头节点，复用原节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0206_reverse_linked_list -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
