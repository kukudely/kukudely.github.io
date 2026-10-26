# 131. 分割回文串

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/palindrome-partitioning/)

**题意：** 把字符串切成若干非空子串，要求每个子串都是回文，返回所有切分方案。题目字符均为小写英文字母。

**示例：** `s = "aab"` → `[["a","a","b"],["aa","b"]]`。

## 方法、入参与返回值

### `partition`

```go
func partition(s string) [][]string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 待分割字符串 |

返回值类型：`[][]string`。使每一段均为回文串的全部分割方案。

## 本地测试约定

额外边界：空字符串有一种空分割方案，返回 [][]string{{}}。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0131_palindrome_partitioning -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
