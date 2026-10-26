# 169. 多数元素

分类：技巧 · 难度：简单

[LeetCode 原题](https://leetcode.cn/problems/majority-element/)

**题意**：数组保证存在一个出现次数大于 n/2 的多数元素，返回该值。

**示例**：nums = [2,2,1,1,1,2,2] → 2。

## 方法、入参与返回值

### `majorityElement`

```go
func majorityElement(nums []int) int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 非空整数数组，保证存在出现次数超过一半的元素 |

返回值类型：`int`。出现次数严格大于 len(nums)/2 的元素。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0169_majority_element -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
