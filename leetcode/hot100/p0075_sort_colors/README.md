# 75. 颜色分类

分类：技巧 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/sort-colors/)

**题意**：数组只包含 0、1、2，原地将相同颜色归拢，结果按 0、1、2 排列。

**示例**：nums = [2,0,2,1,1,0] → [0,0,1,1,2,2]。

## 方法、入参与返回值

### `sortColors`

```go
func sortColors(nums []int)
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 只包含 0、1、2 的数组，需要原地修改 |

返回值类型：`无`。无返回值；将 nums 按 0、1、2 的顺序原地排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0075_sort_colors -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
