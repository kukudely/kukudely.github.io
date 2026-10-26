# 347. 前 K 个高频元素

分类：堆 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/top-k-frequent-elements/)

**题意**：返回出现频率最高的 k 个不同元素，答案顺序任意，题目保证答案集合唯一。

**示例**：nums = [1,1,1,2,2,3], k = 2 → [1,2]。

## 方法、入参与返回值

### `topKFrequent`

```go
func topKFrequent(nums []int, k int) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 整数数组 |
| `k` | `int` | 需要返回的不同元素数量，保证有效 |

返回值类型：`[]int`。出现频次最高的 k 个不同元素，顺序不限。

## 本地测试约定

测试接受不同的返回顺序。若频次并列，只要所选元素满足前 k 高频条件即可。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0347_top_k_frequent_elements -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
