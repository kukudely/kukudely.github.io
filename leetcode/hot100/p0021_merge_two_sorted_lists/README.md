# 21. 合并两个有序链表

分类：链表 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/merge-two-sorted-lists/)

**题意：** 将两个升序单链表合并成一个升序链表，允许复用原节点。

**示例：** list1 = [1,2,4], list2 = [1,3,4] → [1,1,2,3,4,4]。

## 方法、入参与返回值

### `mergeTwoLists`

```go
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `list1` | `*ListNode` | 第一条非递减链表头节点 |
| `list2` | `*ListNode` | 第二条非递减链表头节点 |

返回值类型：`*ListNode`。合并后的非递减链表头节点，复用输入节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0021_merge_two_sorted_lists -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
