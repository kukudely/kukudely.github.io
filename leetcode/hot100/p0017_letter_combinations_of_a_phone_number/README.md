# 17. 电话号码的字母组合

分类：回溯 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/letter-combinations-of-a-phone-number/)

**题意：** 按电话键盘中数字 2–9 的字母映射，返回输入数字串能表示的所有字母组合。

**示例：** `digits = "23"` → `["ad","ae","af","bd","be","bf","cd","ce","cf"]`。

## 方法、入参与返回值

### `letterCombinations`

```go
func letterCombinations(digits string) []string
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `digits` | `string` | 由数字 2 至 9 组成的字符串 |

返回值类型：`[]string`。按电话按键映射得到的全部字母组合；空输入返回空列表。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0017_letter_combinations_of_a_phone_number -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
