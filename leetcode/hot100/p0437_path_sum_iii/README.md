# 437. 路径总和 III

分类：二叉树 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/path-sum-iii/)

**题意：** 统计节点值之和等于 targetSum 的向下路径数量；起点和终点可在任意节点。

**示例：** `root = [10,5,-3,3,2,null,11,3,-2,null,1], targetSum = 8` → `3`。

## 方法、入参与返回值

### `pathSum`

```go
func pathSum(root *TreeNode, targetSum int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `root` | `*TreeNode` | 二叉树根节点 |
| `targetSum` | `int` | 目标路径和 |

返回值类型：`int`。从任意节点向下到后代且节点值之和等于 targetSum 的路径数量。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0437_path_sum_iii -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
