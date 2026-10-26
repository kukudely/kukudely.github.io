# 41. 缺失的第一个正数

分类：普通数组 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/first-missing-positive/)

**题意：** 在未排序整数数组中找出最小缺失正整数。

**示例：** nums = [3,4,-1,1] → 2。

## 方法、入参与返回值

### `firstMissingPositive`

```go
func firstMissingPositive(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 无序整数数组 |

返回值类型：`int`。数组中未出现的最小正整数。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0041_first_missing_positive -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
