# 2. 两数相加

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/add-two-numbers/)

**题意：** 两个非负整数用链表逆序保存，每个节点是一位数字，返回相加后的逆序数字链表。

**示例：** l1 = [2,4,3], l2 = [5,6,4] → [7,0,8]，对应 342+465=807。

## 方法、入参与返回值

### `addTwoNumbers`

```go
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `l1` | `*ListNode` | 第一个非负整数的逆序数字链表 |
| `l2` | `*ListNode` | 第二个非负整数的逆序数字链表 |

返回值类型：`*ListNode`。两个整数之和的逆序数字链表，使用新节点且保留输入链表。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0002_add_two_numbers -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
