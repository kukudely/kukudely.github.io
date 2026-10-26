# 62. 不同路径

分类：多维动态规划 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/unique-paths/)

**题意**：在 m×n 网格中，从左上角走到右下角，每次只能向右或向下，求不同路径数。

**示例**：m = 3, n = 7 → 28。

## 方法、入参与返回值

### `uniquePaths`

```go
func uniquePaths(m int, n int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `m` | `int` | 网格行数，大于 0 |
| `n` | `int` | 网格列数，大于 0 |

返回值类型：`int`。从左上角仅向右或向下移动到右下角的路径数。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0062_unique_paths -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
