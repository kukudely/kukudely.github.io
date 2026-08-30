+++
title = "LeetCode 283：移动零"
date = '2026-08-30T21:33:26+08:00'
draft = false
tags = ["LeetCode", "算法", "数组", "双指针", "Go"]
+++

## 题目描述

[LeetCode 283：移动零](https://leetcode.cn/problems/move-zeroes/description/?envType=study-plan-v2&envId=top-100-liked)

[本次提交记录](https://leetcode.cn/problems/move-zeroes/submissions/745653133/?envType=study-plan-v2&envId=top-100-liked)

给定一个数组 `nums`，将所有 `0` 移动到数组末尾，同时保持非零元素的相对顺序。

要求在不复制数组的情况下原地完成操作。

### 示例 1

```text
输入：nums = [0,1,0,3,12]
输出：[1,3,12,0,0]
```

### 示例 2

```text
输入：nums = [0]
输出：[0]
```

## 题解

### 思路：双指针

使用 `left` 和 `right` 两个指针从左向右遍历数组：

- `right` 负责寻找非零元素。
- `left` 指向下一个非零元素应该放置的位置。
- 当 `nums[right]` 非零时，交换 `nums[left]` 和 `nums[right]`，然后将 `left` 向右移动一位。
- 每轮遍历结束后，将 `right` 向右移动一位。

遍历过程中，`[0, left)` 始终保存已经处理过的非零元素，并保持它们在原数组中的相对顺序。区间 `[left, right)` 保存已经遇到的零。发现新的非零元素时，通过交换把它放到 `left` 指向的位置，同时把零移到后面。

以 `[0,1,0,3,12]` 为例：

```text
初始：       [0,1,0,3,12]
处理 1 后：  [1,0,0,3,12]
处理 3 后：  [1,3,0,0,12]
处理 12 后： [1,3,12,0,0]
```

### 正确性说明

每次处理非零元素时，算法都会把它放到当前最靠左的可用位置。非零元素按照从左到右的顺序依次被处理，因此相对顺序保持不变。

`left` 左侧最终包含全部非零元素，其余位置全部为零，所以结果满足题目要求。

## 复杂度分析

- 时间复杂度：`O(n)`。`right` 只遍历数组一次。
- 空间复杂度：`O(1)`。算法只使用两个指针和一个长度变量，所有操作都在原数组上完成。

## Go 代码

```go
func moveZeroes(nums []int) {
	left, right, n := 0, 0, len(nums)
	for right < n {
		if nums[right] != 0 {
			nums[left], nums[right] = nums[right], nums[left]
			left++
		}
		right++
	}
}
```

## 小结

双指针让查找和放置非零元素在一次遍历中完成。`right` 负责扫描，`left` 负责维护非零区间的边界，最终以 `O(n)` 时间和 `O(1)` 额外空间完成原地移动。
