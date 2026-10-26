# 55. 跳跃游戏

分类：贪心算法 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/jump-game/)

**题意**：从数组下标 0 出发，nums[i] 表示当前位置最多向前跳几步，判断能否到达最后一个下标。

**示例**：nums = [2,3,1,1,4] → true；nums = [3,2,1,0,4] → false。

## 方法、入参与返回值

### `canJump`

```go
func canJump(nums []int) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非负整数数组，每项表示该位置最大向右跳跃长度 |

返回值类型：`bool`。从下标 0 能到达最后一个位置时返回 true。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0055_jump_game -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
