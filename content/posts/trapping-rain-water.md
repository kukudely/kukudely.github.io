+++
title = "LeetCode 42：接雨水"
date = '2026-09-06T14:33:16+08:00'
draft = false
tags = ["LeetCode", "算法", "数组", "动态规划", "双指针", "Go"]
+++

## 题目描述

[LeetCode 42：接雨水](https://leetcode.cn/problems/trapping-rain-water/?envType=study-plan-v2&envId=top-100-liked)

**难度：困难**

给定 `n` 个非负整数表示每个宽度为 `1` 的柱子的高度图，计算下雨之后能够接住的雨水总量。

### 示例 1

```text
输入：height = [0,1,0,2,1,0,1,3,2,1,2,1]
输出：6
```

### 示例 2

```text
输入：height = [4,2,0,3,2,5]
输出：9
```

## 核心思路

对于位置 `i`，能够接住的雨水量取决于它左侧最高柱和右侧最高柱中较矮的一根：

```text
water[i] = min(leftMax[i], rightMax[i]) - height[i]
```

数组两端缺少一侧边界，无法接住雨水。只需计算中间位置的积水量，再将它们累加起来。

## 题解一：动态规划

### 思路

创建两个长度为 `n` 的数组：

- `leftMax[i]` 表示从下标 `0` 到 `i` 的最大高度。
- `rightMax[i]` 表示从下标 `i` 到 `n-1` 的最大高度。

先从左向右计算 `leftMax`，再从右向左计算 `rightMax`。得到两侧最高柱后，可以直接计算每个位置的积水量。

以位置 `i` 为例，水面高度受到左右两侧较矮边界的限制，因此水面高度为 `min(leftMax[i], rightMax[i])`。减去当前位置的柱高，就是该位置能够接住的雨水量。

### 正确性说明

`leftMax[i]` 和 `rightMax[i]` 分别记录位置 `i` 两侧能够形成边界的最高柱。水无法越过较矮的一侧，所以位置 `i` 的最高水面由两者的较小值确定。

每个位置的计算结果都准确表示该列上方的积水量。累加所有位置后，即可得到总雨水量。

### 复杂度分析

- 时间复杂度：`O(n)`。需要三次线性遍历。
- 空间复杂度：`O(n)`。两个辅助数组各保存 `n` 个元素。

### Go 代码

```go
func trapDP(height []int) int {
	n := len(height)
	if n == 0 {
		return 0
	}

	leftMax := make([]int, n)
	rightMax := make([]int, n)

	leftMax[0] = height[0]
	for i := 1; i < n; i++ {
		leftMax[i] = max(leftMax[i-1], height[i])
	}

	rightMax[n-1] = height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax[i] = max(rightMax[i+1], height[i])
	}

	ans := 0
	for i := 0; i < n; i++ {
		ans += min(leftMax[i], rightMax[i]) - height[i]
	}

	return ans
}
```

## 题解二：双指针

### 思路

动态规划保存了每个位置的左右最大高度。继续观察可以发现，计算当前积水量时只需要已经遍历过的 `leftMax` 和 `rightMax`，因此可以使用双指针压缩空间。

设置 `left`、`right` 两个指针，并维护：

- `leftMax`：从左端到 `left` 的最大高度。
- `rightMax`：从右端到 `right` 的最大高度。

当 `height[left] < height[right]` 时，右侧至少存在一根比当前左柱更高的柱子，左侧位置的积水量可以由 `leftMax` 确定。处理 `left` 后向右移动。

其余情况下，左侧边界足以支撑当前右侧位置，积水量可以由 `rightMax` 确定。处理 `right` 后向左移动。

### 正确性说明

双指针始终优先处理两端高度较小的一侧。对于较小的一侧，另一侧已经提供了足够高的边界，因此当前积水量只受本侧历史最高柱限制。

处理完成后移动对应指针，不会影响已经确定的积水量。左右指针逐步收拢，每个位置都会被准确处理一次，累加结果就是总雨水量。

### 复杂度分析

- 时间复杂度：`O(n)`。左右指针合计遍历整个数组一次。
- 空间复杂度：`O(1)`。只使用常数个变量。

### Go 代码

```go
func trap(height []int) (ans int) {
	left, right := 0, len(height)-1
	leftMax, rightMax := 0, 0
	for left < right {
		leftMax = max(leftMax, height[left])
		rightMax = max(rightMax, height[right])
		if height[left] < height[right] {
			ans += leftMax - height[left]
			left++
		} else {
			ans += rightMax - height[right]
			right--
		}
	}
	return
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
```

## 两种解法对比

| 解法 | 时间复杂度 | 空间复杂度 | 特点 |
|---|---:|---:|---|
| 动态规划 | `O(n)` | `O(n)` | 思路直观，显式保存每个位置两侧的最大高度 |
| 双指针 | `O(n)` | `O(1)` | 维护两侧最大高度，省去辅助数组 |

## 小结

这道题的关键是确定每个位置的水面上限。动态规划完整记录左右两侧的最高柱，便于理解每列雨水的来源；双指针利用两侧边界逐步确定结果，将额外空间降到 `O(1)`。
