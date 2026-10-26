# 295. 数据流的中位数

分类：堆 · 难度：困难

[LeetCode 原题](https://leetcode.cn/problems/find-median-from-data-stream/)

**题意**：设计数据结构，不断接收整数，并能查询当前所有元素的中位数。

**示例**：AddNum(1)、AddNum(2) 后 FindMedian() = 1.5；再 AddNum(3) 后为 2。

## 方法、入参与返回值

### `Constructor`

```go
func Constructor() MedianFinder
```

入参：无。

返回值类型：`MedianFinder`。新建空 MedianFinder。

### `AddNum`

```go
func (f *MedianFinder) AddNum(num int)
```

接收者：`*MedianFinder`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `num` | `int` | 需要加入数据流的整数 |

返回值类型：`无`。无返回值；保存新元素。

### `FindMedian`

```go
func (f *MedianFinder) FindMedian() float64
```

接收者：`*MedianFinder`，操作同一个实例的状态。

入参：无。

返回值类型：`float64`。当前全部元素的中位数，偶数个时取中间两数平均；调用时保证非空。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0295_find_median_from_data_stream -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
