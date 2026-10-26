# 45. 跳跃游戏 II

分类：贪心算法 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/jump-game-ii/)

**题意**：从下标 0 出发，nums[i] 为最大前跳距离，求到达最后一个位置的最少跳跃次数。

**示例**：nums = [2,3,1,1,4] → 2。

## 方法、入参与返回值

### `jump`

```go
func jump(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非负整数数组，每项表示最大跳跃长度，原题保证可以到达终点 |

返回值类型：`int`。从下标 0 到最后一个位置的最少跳跃次数。

## 本地测试约定

额外边界约定：空数组返回 0；本题测试的非空输入均可到达终点。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0045_jump_game_ii -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
