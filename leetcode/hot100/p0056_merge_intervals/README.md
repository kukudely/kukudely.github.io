# 56. 合并区间

分类：普通数组 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/merge-intervals/)

**题意：** 合并所有相交或相接的闭区间，使输出区间覆盖相同范围且互不重叠。

**示例：** intervals = [[1,3],[2,6],[8,10],[15,18]] → [[1,6],[8,10],[15,18]]。

## 方法、入参与返回值

### `merge`

```go
func merge(intervals [][]int) [][]int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `intervals` | `[][]int` | 闭区间列表，每项为 [起点,终点] |

返回值类型：`[][]int`。合并全部重叠区间后的区间列表，按区间起点升序排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0056_merge_intervals -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
