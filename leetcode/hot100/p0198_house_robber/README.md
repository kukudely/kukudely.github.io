# 198. 打家劫舍

分类：动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/house-robber/)

**题意**：房屋首尾不相接，不能偷相邻房屋，求可以取得的最大金额。

**示例**：nums = [2,7,9,3,1] → 12，选择 2、9、1。

## 方法、入参与返回值

### `rob`

```go
func rob(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 按房屋顺序给出的非负金额 |

返回值类型：`int`。不选择相邻房屋时能够取得的最大金额；空数组返回 0。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0198_house_robber -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
