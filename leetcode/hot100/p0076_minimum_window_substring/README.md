# 76. 最小覆盖子串

分类：子串 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/minimum-window-substring/)

**题意：** 在 s 中找到包含 t 全部字符及其重复次数的最短连续子串；无解返回空串。题目输入由英文字母组成。

**示例：** s = "ADOBECODEBANC", t = "ABC" → "BANC"。

## 方法、入参与返回值

### `minWindow`

```go
func minWindow(s string, t string) string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `s` | `string` | 待截取字符串 |
| `t` | `string` | 需要覆盖的目标字符串，重复字符也要满足次数 |

返回值类型：`string`。包含 t 全部字符的最短连续子串；不存在时返回空字符串。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0076_minimum_window_substring -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
