# 49. 字母异位词分组

分类：哈希 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/group-anagrams/)

**题意：** 把字符串按字母异位词关系分组；组间和组内顺序均不限。题目字符串只含小写英文字母。

**示例：** strs = ["eat","tea","tan","ate","nat","bat"] → [["eat","tea","ate"],["tan","nat"],["bat"]]。

## 方法、入参与返回值

### `groupAnagrams`

```go
func groupAnagrams(strs []string) [][]string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `strs` | `[]string` | 仅包含小写英文字母的字符串数组 |

返回值类型：`[][]string`。按字母异位词分组后的字符串二维数组；组间和组内顺序不限。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0049_group_anagrams -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
