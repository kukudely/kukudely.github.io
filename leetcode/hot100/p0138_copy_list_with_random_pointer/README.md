# 138. 随机链表的复制

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/copy-list-with-random-pointer/)

**题意：** 复制包含 Next 和 Random 的链表，所有新指针必须指向复制出来的节点或 nil。

**示例：** 原链用 [值,random下标] 表示为 [[7,null],[13,0],[11,1]]；复制链的表示相同，所有节点地址均不同。

## 方法、入参与返回值

### `copyRandomList`

```go
func copyRandomList(head *Node) *Node
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `head` | `*Node` | 带 Next 和 Random 指针的链表头节点 |

返回值类型：`*Node`。深拷贝链表的头节点；复制节点与原节点完全独立，并保持随机指向关系。

## 本地测试约定

复制过程中可以临时调整原链连接；返回前须恢复。Next/Random 的比较会检查指针身份，无法用仅相同的节点值代替。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0138_copy_list_with_random_pointer -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
