# 763. 划分字母区间

分类：贪心算法 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/partition-labels/)

**题意**：把小写字母字符串分成尽量多的连续片段，使每个字母只出现在一个片段中，返回各片段长度。

**示例**：s = "ababcbacadefegdehijhklij" → [9,7,8]。

## 方法、入参与返回值

### `partitionLabels`

```go
func partitionLabels(s string) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 仅含小写英文字母的字符串 |

返回值类型：`[]int`。尽量多分段且同一字母仅出现在一个片段时，各片段的长度。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0763_partition_labels -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
