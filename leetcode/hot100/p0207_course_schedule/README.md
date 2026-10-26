# 207. 课程表

分类：图论 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/course-schedule/)

**题意：** 课程对 [a,b] 表示先学 b 才能学 a，判断所有课程能否完成，即依赖有向图是否无环。

**示例：** `numCourses = 2, prerequisites = [[1,0]]` → `true`；加上 `[0,1]` 后为 `false`。

## 方法、入参与返回值

### `canFinish`

```go
func canFinish(numCourses int, prerequisites [][]int) bool
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `numCourses` | `int` | 课程总数，课程编号为 0 至 numCourses-1 |
| `prerequisites` | `[][]int` | 依赖列表，每项 [a,b] 表示学 a 前先学 b |

返回值类型：`bool`。能够完成全部课程时为 true，存在依赖环时为 false。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0207_course_schedule -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
