# 1. 两数之和

分类：哈希 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/two-sum/)

**题意：** 在整数数组中找两个不同位置，使对应元素之和等于 target，返回下标。题目保证恰有一组答案。

**示例：** nums = [2,7,11,15], target = 9 → [0,1]。

## 方法、入参与返回值

### `twoSum`

```go
func twoSum(nums []int, target int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |
| `target` | `int` | 目标和 |

返回值类型：`[]int`。两个不同元素的原数组下标，顺序不限；输入保证存在唯一答案。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0001_two_sum -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
