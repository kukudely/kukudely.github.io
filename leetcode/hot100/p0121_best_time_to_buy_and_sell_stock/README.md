# 121. 买卖股票的最佳时机

分类：贪心算法 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/best-time-to-buy-and-sell-stock/)

**题意**：最多进行一次买入和一次卖出，卖出日期必须晚于买入，求最大利润；无利可图时返回 0。

**示例**：prices = [7,1,5,3,6,4] → 5。

## 方法、入参与返回值

### `maxProfit`

```go
func maxProfit(prices []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `prices` | `[]int` | 每日股票价格 |

返回值类型：`int`。最多完成一次先买后卖所能获得的最大利润；不交易时利润为 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0121_best_time_to_buy_and_sell_stock -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
