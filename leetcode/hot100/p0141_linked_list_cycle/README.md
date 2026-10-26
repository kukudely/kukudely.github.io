# 141. 环形链表

分类：链表 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/linked-list-cycle/)

**题意：** 判断沿 Next 指针前进是否会再次访问同一个节点。

**示例：** head = [3,2,0,-4]，尾节点指向下标1的节点 → true。

## 方法、入参与返回值

### `hasCycle`

```go
func hasCycle(head *ListNode) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点，链表可能存在环 |

返回值类型：`bool`。存在环返回 true，否则返回 false。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0141_linked_list_cycle -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
