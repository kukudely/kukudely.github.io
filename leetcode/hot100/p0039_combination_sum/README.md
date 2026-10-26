# 39. 组合总和

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/combination-sum/)

**题意：** 给定互不相同的正整数 candidates，返回和为 target 的所有组合。每个数字可重复使用，组合中顺序不同视为同一种。

**示例：** `candidates = [2,3,6,7], target = 7` → `[[2,2,3],[7]]`。

## 方法、入参与返回值

### `combinationSum`

```go
func combinationSum(candidates []int, target int) [][]int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `candidates` | `[]int` | 互不相同的正整数候选值，每个值可以重复使用 |
| `target` | `int` | 需要组成的目标和 |

返回值类型：`[][]int`。所有不重复的组合；同一组合中元素的选取顺序不产生新答案。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0039_combination_sum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
