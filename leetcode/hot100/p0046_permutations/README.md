# 46. 全排列

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/permutations/)

**题意：** 给定互不相同的整数，返回它们的所有排列，答案顺序不限。

**示例：** `nums = [1,2,3]` → `[[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]`。

## 方法、入参与返回值

### `permute`

```go
func permute(nums []int) [][]int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 元素互不相同的整数数组 |

返回值类型：`[][]int`。全部排列，排列之间的顺序不限。

## 本地测试约定

额外边界：空输入返回一个空排列，即 [][]int{{}}。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0046_permutations -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
