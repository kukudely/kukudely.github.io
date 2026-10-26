# 322. 零钱兑换

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/coin-change/)

**题意**：硬币面额均为正整数，每种可无限使用。求组成 amount 的最少硬币数，无法组成时返回 -1。

**示例**：coins = [1,2,5], amount = 11 → 3；coins = [2], amount = 3 → -1。

## 方法、入参与返回值

### `coinChange`

```go
func coinChange(coins []int, amount int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `coins` | `[]int` | 可重复使用的正整数硬币面额 |
| `amount` | `int` | 非负目标金额 |

返回值类型：`int`。组成 amount 所需的最少硬币数；无法组成时为 -1，金额 0 返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0322_coin_change -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
