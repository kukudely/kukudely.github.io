# 70. 爬楼梯

分类：动态规划 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/climbing-stairs/)

**题意**：每次可以爬 1 或 2 个台阶，求爬到第 n 阶的不同方案数。

**示例**：n = 3 → 3；三种方案为 1+1+1、1+2、2+1。

## 方法、入参与返回值

### `climbStairs`

```go
func climbStairs(n int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `n` | `int` | 台阶总数，每步允许爬 1 或 2 阶 |

返回值类型：`int`。爬到第 n 阶的不同方案数量。测试约定 n=0 时为空方案，返回 1。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0070_climbing_stairs -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
