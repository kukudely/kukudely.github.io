# 146. LRU 缓存

分类：链表 · 难度：中等

[LeetCode 原题](https://leetcode.cn/problems/lru-cache/)

**题意：** 实现容量固定的 LRU 缓存：Get 读取并刷新最近使用顺序；Put 写入或更新，并在超容量时淘汰最久未使用键。

**示例：** 容量2：Put(1,1), Put(2,2), Get(1)=1, Put(3,3), Get(2)=-1。

## 方法、入参与返回值

### `Constructor`

```go
func Constructor(capacity int) LRUCache
```

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `capacity` | `int` | 缓存最大容量 |

返回值类型：`LRUCache`。新建 LRUCache。容量为 0 时 Put 不保存元素。

### `Get`

```go
func (c *LRUCache) Get(key int) int
```

接收者：`*LRUCache`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `key` | `int` | 待查询键 |

返回值类型：`int`。命中时返回值并更新最近使用顺序；未命中返回 -1。

### `Put`

```go
func (c *LRUCache) Put(key int, value int)
```

接收者：`*LRUCache`，操作同一个实例的状态。

| 入参 | Go 类型 | 含义 |
|---|---|---|
| `key` | `int` | 需要写入的键 |
| `value` | `int` | 需要写入的值 |

返回值类型：`无`。无返回值；写入或更新后标记为最近使用，超容量时淘汰最久未使用的键。

## 本地测试约定

原题要求 Get、Put 平均 O(1)；本地功能测试验证语义，不通过执行速度推断复杂度。

## 开始练习

1. 在 `solution.go` 中实现上述方法，并移除相应的 `panic("TODO")`。
2. 如有 `types.go`，节点字段已定义；设计题的结构体字段由你补充。
3. 阅读 `solution_test.go` 中的输入、期望输出和断言，运行测试。

在 `leetcode/` 目录执行：

```sh
go test ./hot100/p0146_lru_cache -v -count=1 -timeout=30s
```

未实现方法显示 `SKIP`，完成后检查测试是否显示 `PASS`。异常 panic 和不匹配的结果会使测试失败。每题独立一个包，可以直接使用 LeetCode 的函数名。

[返回题目目录](../../README.md)
