+++
title = "LeetCode 128：最长连续序列"
date = '2026-08-30T00:00:00+08:00'
draft = true
tags = ["LeetCode", "算法", "哈希表", "Go"]
+++

## 题目描述

[LeetCode 128：最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/description/?envType=study-plan-v2&envId=top-100-like)

给定一个未排序的整数数组 `nums`，找出数字连续的最长序列的长度。序列中的元素不要求在原数组中连续。

要求设计并实现时间复杂度为 `O(n)` 的算法。

### 示例 1

```text
输入：nums = [100,4,200,1,3,2]
输出：4
解释：最长连续序列是 [1,2,3,4]，长度为 4。
```

### 示例 2

```text
输入：nums = [0,3,7,2,5,8,4,6,0,1]
输出：9
```

### 示例 3

```text
输入：nums = [1,0,1,2]
输出：3
```

## 题解

### 思路：哈希集合

先把数组中的所有数字放入哈希集合。这样既能去重，也能在平均 `O(1)` 时间内判断某个数字是否存在。

遍历集合中的每个数字 `num`：

1. 如果 `num - 1` 也在集合中，说明 `num` 已有前驱，直接跳过。
2. 如果 `num - 1` 不在集合中，说明 `num` 是一段连续序列的起点。
3. 从 `num` 开始不断检查 `num + 1`、`num + 2`……是否存在，统计当前连续序列的长度。
4. 用当前长度更新最长长度。

关键点是：**只从连续序列的起点向后查找。**

例如集合中有 `1、2、3、4`。遍历到 `2、3、4` 时，因为它们都有前驱，所以不会重复向后查找；只有 `1` 会完整扫描这段序列。

### 正确性说明

任意一段连续序列都有且只有一个最小值。这个最小值的前一个数字一定不在集合中，因此算法一定会把它识别为序列起点，并从该起点找到整段连续序列。

其他数字因为存在前驱，不会被当作起点，所以不会重复统计同一段序列。遍历所有起点并取最大长度，得到的就是最长连续序列的长度。

## 复杂度分析

- 时间复杂度：`O(n)`。构建哈希集合需要 `O(n)`；虽然存在内层循环，但每个数字只会在所属连续序列从起点扩展时被访问一次，因此总扩展次数不超过 `n`。
- 空间复杂度：`O(n)`。哈希集合最多保存 `n` 个不同的数字。

## Go 代码

```go
func longestConsecutive(nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, num := range nums {
		set[num] = struct{}{}
	}

	longest := 0

	for num := range set {
		// num - 1 存在时，num 已有前驱，直接跳过。
		if _, exists := set[num-1]; exists {
			continue
		}

		current := num
		length := 1

		for {
			if _, exists := set[current+1]; !exists {
				break
			}
			current++
			length++
		}

		if length > longest {
			longest = length
		}
	}

	return longest
}
```

## 小结

这道题的重点是避免从每个数字开始重复扫描。哈希表提供平均 `O(1)` 的查询效率；判断 `num - 1` 是否存在，可以快速筛出每段连续序列唯一的起点，从而将整体时间复杂度控制在 `O(n)`。
