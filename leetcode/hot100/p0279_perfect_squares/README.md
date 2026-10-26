# 279. 完全平方数

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/perfect-squares/)

**题意**：把正整数 n 写成若干完全平方数之和，返回最少使用多少个平方数。

**示例**：n = 12 → 3（4+4+4）；n = 13 → 2（4+9）。

## 方法、入参与返回值

### `numSquares`

```go
func numSquares(n int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `n` | `int` | 非负目标整数 |

返回值类型：`int`。组成 n 所需的最少完全平方数个数，平方数可重复使用；n=0 返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0279_perfect_squares -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
