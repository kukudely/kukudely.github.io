# 23. 合并 K 个升序链表

分类：链表 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/merge-k-sorted-lists/)

**题意：** 合并 k 条升序单链表，输出一条升序链表。

**示例：** lists = [[1,4,5],[1,3,4],[2,6]] → [1,1,2,3,4,4,5,6]。

## 方法、入参与返回值

### `mergeKLists`

```go
func mergeKLists(lists []*ListNode) *ListNode
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `lists` | `[]*ListNode` | 多个非递减链表的头节点列表，允许 nil |

返回值类型：`*ListNode`。将全部链表合并后的非递减链表头节点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0023_merge_k_sorted_lists -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
