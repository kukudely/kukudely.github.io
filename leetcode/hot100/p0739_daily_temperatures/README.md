# 739. 每日温度

分类：栈 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/daily-temperatures/)

**题意**：对每天的温度，返回还需等待多少天才会出现严格更高的温度；未来没有更高温度时填 0。

**示例**：temperatures = [73,74,75,71,69,72,76,73] → [1,1,4,2,1,1,0,0]。

## 方法、入参与返回值

### `dailyTemperatures`

```go
func dailyTemperatures(temperatures []int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `temperatures` | `[]int` | 每日温度列表 |

返回值类型：`[]int`。每一天距离下一次更高温度的天数；不存在更高温度时为 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0739_daily_temperatures -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
