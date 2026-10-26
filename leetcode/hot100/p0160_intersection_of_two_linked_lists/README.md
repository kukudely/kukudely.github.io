# 160. 相交链表

分类：链表 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/intersection-of-two-linked-lists/)

**题意：** 两个无环单链表若共享尾部节点，返回第一个共享节点；否则返回 nil。

**示例：** A = [4,1] + 公共链 [8,4,5]，B = [5,6,1] + 同一公共链 → 返回值为8的公共节点。

## 方法、入参与返回值

### `getIntersectionNode`

```go
func getIntersectionNode(headA, headB *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `headA` | `*ListNode` | 第一条无环单链表的头节点 |
| `headB` | `*ListNode` | 第二条无环单链表的头节点 |

返回值类型：`*ListNode`。两条链表相交的首个节点指针；不相交返回 nil。比较节点身份。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0160_intersection_of_two_linked_lists -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
