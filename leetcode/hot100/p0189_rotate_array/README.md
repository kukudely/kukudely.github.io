# 189. 轮转数组

分类：普通数组 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/rotate-array/)

**题意：** 将数组原地向右轮转 k 步，k 为非负整数。

**示例：** nums = [1,2,3,4,5,6,7], k = 3 → [5,6,7,1,2,3]。

## 方法、入参与返回值

### `rotate`

```go
func rotate(nums []int, k int)
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `nums` | `[]int` | 需要原地修改的整数数组 |
| `k` | `int` | 向右轮转的非负步数 |

返回值类型：`无`。无返回值；直接将 nums 向右轮转 k 步。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0189_rotate_array -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
