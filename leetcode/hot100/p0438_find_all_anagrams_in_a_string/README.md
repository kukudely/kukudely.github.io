# 438. 找到字符串中所有字母异位词

分类：滑动窗口 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/find-all-anagrams-in-a-string/)

**题意：** 返回 s 中所有长度等于 p、且字符计数与 p 相同的子串起点。输入只含小写英文字母。

**示例：** s = "cbaebabacd", p = "abc" → [0,6]。

## 方法、入参与返回值

### `findAnagrams`

```go
func findAnagrams(s string, p string) []int
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 待搜索字符串 |
| `p` | `string` | 目标模式串 |

返回值类型：`[]int`。s 中所有与 p 互为异位词的子串起始下标，按从小到大排列。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0438_find_all_anagrams_in_a_string -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
