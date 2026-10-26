# 142. 环形链表 II

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/linked-list-cycle-ii/)

**题意：** 若链表有环，返回入环的第一个节点；无环返回 nil。

**示例：** head = [3,2,0,-4]，尾节点指向下标1 → 返回下标1、值为2的原节点。

## 方法、入参与返回值

### `detectCycle`

```go
func detectCycle(head *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点，链表可能存在环 |

返回值类型：`*ListNode`。入环节点的原始指针；无环时返回 nil。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0142_linked_list_cycle_ii -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
