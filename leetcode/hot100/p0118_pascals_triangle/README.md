# 118. 杨辉三角

分类：动态规划 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/pascals-triangle/)

**题意**：返回杨辉三角的前 numRows 行。每行两端为 1，内部数字等于上一行相邻两个数字之和。

**示例**：numRows = 5 → [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]。

## 方法、入参与返回值

### `generate`

```go
func generate(numRows int) [][]int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `numRows` | `int` | 需要生成的行数 |

返回值类型：`[][]int`。杨辉三角前 numRows 行，第 0 行为 [1]。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0118_pascals_triangle -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
