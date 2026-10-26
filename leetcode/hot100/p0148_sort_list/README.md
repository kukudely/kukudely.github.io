# 148. 排序链表

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/sort-list/)

**题意：** 将单链表按节点值升序排列，返回排序后的头节点。

**示例：** head = [4,2,1,3] → [1,2,3,4]。

## 方法、入参与返回值

### `sortList`

```go
func sortList(head *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*ListNode` | 单链表头节点 |

返回值类型：`*ListNode`。按节点值非递减排列的链表头节点，复用原节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0148_sort_list -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
