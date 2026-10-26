# 155. 最小栈

分类：栈 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/min-stack/)

**题意**：设计栈，支持 Push、Pop、Top、GetMin，并要求每次操作为 O(1)。

**示例**：依次 Push(-2)、Push(0)、Push(-3)，GetMin() 为 -3；Pop() 后 Top() 为 0，GetMin() 为 -2。

## 方法、入参与返回值

### `Constructor`

```go
func Constructor() MinStack
```

入参：无。

返回值类型：`MinStack`。新建空 MinStack。

### `Push`

```go
func (s *MinStack) Push(val int)
```

接收者：`*MinStack`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `val` | `int` | 待入栈整数 |

返回值类型：`无`。无返回值；将 val 压入栈顶。

### `Pop`

```go
func (s *MinStack) Pop()
```

接收者：`*MinStack`，操作同一个实例的状态。

入参：无。

返回值类型：`无`。无返回值；删除栈顶，调用时保证栈非空。

### `Top`

```go
func (s *MinStack) Top() int
```

接收者：`*MinStack`，操作同一个实例的状态。

入参：无。

返回值类型：`int`。当前栈顶值，调用时保证栈非空。

### `GetMin`

```go
func (s *MinStack) GetMin() int
```

接收者：`*MinStack`，操作同一个实例的状态。

入参：无。

返回值类型：`int`。当前栈中的最小值，调用时保证栈非空。

## 本地测试约定

原题要求每个操作 O(1)；本地功能测试验证栈和最小值语义。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0155_min_stack -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
