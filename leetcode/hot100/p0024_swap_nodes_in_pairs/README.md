# 24. 两两交换链表中的节点

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/swap-nodes-in-pairs/)

**题意：** 每两个相邻节点交换位置，节点值保持不变；奇数长度的最后一个节点保持原位。

**示例：** head = [1,2,3,4] → [2,1,4,3]。

## 方法、入参与返回值

### `swapPairs`

```go
func swapPairs(head *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点 |

返回值类型：`*ListNode`。两两交换实际节点后的头节点；最后落单节点保留。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0024_swap_nodes_in_pairs -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
