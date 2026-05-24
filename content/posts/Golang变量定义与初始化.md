+++
title = "Golang变量定义与初始化"
date = '2026-05-24T13:51:50+0800'
draft = false
tags = ["Golang", "基础语法", "算法"]
+++



## 一、Go 中变量的几种定义方式

### 1. `var 变量名 类型`

只定义，不手动赋值，变量会拿到该类型的零值。

```go
var a int
var s string
var ok bool
```

等价于：

```go
var a int = 0
var s string = ""
var ok bool = false
```

为什么这样定义：

- 适合先声明，后赋值
- 适合需要明确类型的场景
- 适合结果变量、函数外变量、包级变量

最常见场景：

- `var ans int`
- `var res []int`
- `var root *TreeNode`

### 2. `var 变量名 类型 = 值`

定义时显式指定类型，并立即初始化。

```go
var n int = 10
var name string = "go"
```

为什么这样定义：

- 想明确表达“我就要这个类型”
- 避免类型推导带来的歧义
- 可读性比纯推导更强

最常见场景：

- 需要固定整数或浮点类型
- 包级变量初始化

### 3. `var 变量名 = 值`

定义并初始化，类型由编译器自动推导。

```go
var n = 10
var s = "hello"
```

为什么这样定义：

- 比显式写类型更简洁
- 仍然可以用于函数外
- 当类型很明显时，不需要重复写类型名

### 4. `变量名 := 值`

短变量声明，只能在函数内部使用。

```go
n := 10
s := "hello"
nums := []int{1, 2, 3}
```

为什么这样定义：

- 最简洁
- 局部变量最常用
- 刷题时写起来最快

最常见场景：

- 下标
- 计数器
- 临时变量
- 局部数组、切片、map

注意：

- 不能在函数外使用
- 左边至少要有一个新变量

### 5. 多变量定义

```go
var a, b int
var x, y = 1, 2
i, j := 3, 4
```

为什么这样定义：

- 一组相关变量同时声明时更简洁
- 左右边逻辑更对称

## 二、为什么 Go 强调“零值”

Go 中变量即使不手动赋值，也一定会拿到默认值。

| 类型 | 零值 |
|---|---|
| `int` | `0` |
| `float64` | `0` |
| `bool` | `false` |
| `string` | `""` |
| 指针 | `nil` |
| `slice` | `nil` |
| `map` | `nil` |
| `channel` | `nil` |

为什么重要：

- Go 不允许“未初始化但被使用”
- 刷题中可以直接利用零值
- 例如计数器默认是 `0`，布尔标记默认是 `false`

## 三、不同类型最常用的定义方式

### 1. `int`

最常用：

```go
n := 10
count := 0
```

为什么：

- 局部变量中最简洁
- 类型推导非常明确
- 刷题里的下标、边界、计数器都适合这样写

如果是先声明结果变量：

```go
var ans int
```

为什么：

- 默认就是 `0`
- 后续更新自然

### 2. `string`

最常用：

```go
s := "abc"
var res string
```

为什么：

- 有初始值时直接推导
- 没初始值时零值 `""` 很适合作结果变量

### 3. `bool`

最常用：

```go
found := false
var visited bool
```

为什么：

- `bool` 零值就是 `false`
- 非常适合做标记位

### 4. 数组 `array`

最常用：

```go
arr := [3]int{1, 2, 3}
var nums [5]int
```

为什么：

- 长度固定，初始化直观
- 但数组不够灵活，所以刷题里远没有切片常用

### 5. 切片 `slice`

最常用有两种。

#### 字面量方式

```go
nums := []int{1, 2, 3}
```

为什么这样定义：

- 已知初始内容时最直观
- 测试数据、小型初始集合最适合

#### `make` 方式

```go
dp := make([]int, n)
path := make([]int, 0, n)
```

为什么这样定义：

- `slice` 本质上依赖底层数组
- `make` 会创建底层可用存储
- 长度和容量一眼可见
- 对 DP、辅助数组、回溯路径最实用

这是切片最常用的定义方式之一，原因是：

1. 创建后直接可用
2. 可明确控制长度和容量
3. 更符合刷题代码习惯
4. 对性能更友好

### 6. `map`

最常用：

```go
cnt := make(map[int]int)
visited := make(map[int]bool)
group := make(map[string]int)
```

为什么 `map` 最常用 `make`：

因为只有声明：

```go
var m map[int]int
```

这时 `m` 是 `nil`，不能直接写入：

```go
m[1] = 2 // panic
```

而 `make` 后才真正可用：

```go
m := make(map[int]int)
```

为什么最常这样写：

1. 创建后能直接读写
2. 语义明确
3. 刷题里哈希表几乎都这么定义

### 7. 结构体 `struct`

最常用：

```go
type Pair struct {
    X int
    Y int
}

p := Pair{X: 1, Y: 2}
var p2 Pair
```

为什么最常用字段名初始化：

- 可读性好
- 不容易写错位置
- 字段很多时更清晰

为什么有时用 `var p2 Pair`：

- 字段默认零值
- 适合后续逐个赋值

### 8. 结构体指针

最常用：

```go
node := &ListNode{Val: 1}
dummy := &ListNode{}
root := &TreeNode{Val: 3}
```

为什么最常这样定义：

1. 一行完成创建和初始化
2. 可读性好
3. 链表、树题里最自然

对比：

```go
node := new(ListNode)
node.Val = 1
```

虽然也合法，但不如 `&ListNode{}` 直接、直观。

### 9. 指针

普通变量指针最常见写法：

```go
x := 10
p := &x
```

为什么：

- 最直观
- Go 刷题里更常见的是结构体指针，而不是普通变量指针

## 四、`make` 和 `new` 的区别

### `make`

用于：

- `slice`
- `map`
- `channel`

```go
s := make([]int, 0, 10)
m := make(map[int]int)
```

作用：

- 初始化底层数据结构
- 返回的是该类型本身，不是指针

### `new`

用于任意类型：

```go
p := new(int)
```

返回值是 `*int`。

作用：

- 分配内存
- 返回指针
- 内容为零值

### 为什么 `slice / map / channel` 更常用 `make`

因为这些类型需要的是“可用的内部结构”，而不是“指向零值的指针”。

例如：

```go
var m map[int]int
// m 为 nil，不能写入

m2 := make(map[int]int)
// m2 可直接使用
```

## 五、刷题时推荐的变量定义风格

### 普通局部变量

```go
n := 0
left, right := 0, len(nums)-1
```

### 结果变量

```go
var ans int
var res []int
```

### 切片

```go
nums := []int{1, 2, 3}
dp := make([]int, n)
path := make([]int, 0, n)
```

### map

```go
cnt := make(map[int]int)
visited := make(map[int]bool)
```

### 结构体 / 节点

```go
dummy := &ListNode{}
node := &TreeNode{Val: 1}
```
