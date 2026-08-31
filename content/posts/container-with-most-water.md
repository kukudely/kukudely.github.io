+++
title = "LeetCode 11：盛最多水的容器"
date = '2026-08-31T10:11:25+08:00'
draft = false
tags = ["LeetCode", "算法", "数组", "双指针", "贪心", "Go"]
+++

## 题目描述

[LeetCode 11：盛最多水的容器](https://leetcode.cn/problems/container-with-most-water/description/?envType=study-plan-v2&envId=top-100-liked)

**难度：中等**

给定一个长度为 `n` 的整数数组 `height`。数组中的第 `i` 个元素表示一条垂线的高度，这条垂线的两个端点分别是 `(i, 0)` 和 `(i, height[i])`。

从中选择两条垂线，使它们与 `x` 轴共同构成一个容器，返回该容器可以储存的最大水量。容器不能倾斜。

### 示例 1

```text
输入：height = [1,8,6,2,5,4,8,3,7]
输出：49
解释：选择下标 1 和 8 对应的两条垂线，高度分别为 8 和 7。
      容器宽度为 8 - 1 = 7，有效高度为 min(8, 7) = 7，
      因此水量为 7 × 7 = 49。
```

### 示例 2

```text
输入：height = [1,1]
输出：1
```

## 题解

### 面积如何计算

选择下标 `left` 和 `right` 处的两条垂线后：

- 容器宽度为 `right - left`。
- 水面高度由较短的垂线决定，即 `min(height[left], height[right])`。

因此当前容器的水量为：

```text
(right - left) × min(height[left], height[right])
```

### 思路：双指针

使用 `left` 和 `right` 两个指针，分别从数组的左右两端开始：

1. 根据当前宽度和两侧较短的高度计算水量，并更新最大值。
2. 比较 `height[left]` 和 `height[right]`。
3. 将较短一侧的指针向中间移动一位。
4. 重复以上过程，直到两个指针相遇。

初始状态下，两条垂线之间的距离最大。随着指针向中间移动，容器宽度一定会减小。想要获得更大的水量，有效高度必须增加，因此每次都移动当前较短的一侧，尝试寻找更高的垂线。

以 `[1,8,6,2,5,4,8,3,7]` 为例：

```text
left  right  两侧高度  当前水量  操作
0     8      1 和 7    8         左侧较短，left 右移
1     8      8 和 7    49        右侧较短，right 左移
1     7      8 和 3    18        右侧较短，right 左移
1     6      8 和 8    40        高度相等，移动任意一侧
```

后续区间的宽度继续减小，最终得到最大水量 `49`。

### 为什么移动较短的一侧

假设 `height[left] <= height[right]`，当前水量为：

```text
(right - left) × height[left]
```

如果保留 `left` 并将 `right` 向左移动，新容器的宽度会变小，有效高度仍然不会超过 `height[left]`。这个新容器的水量一定小于当前水量，所以保留当前较短边无法得到更优解。

将 `left` 向右移动后，虽然宽度也会减小，却有机会找到一条更高的垂线，从而提高有效高度。`height[right] < height[left]` 时同理，应当移动 `right`。

### 正确性说明

每轮计算当前两条垂线形成的水量后，算法都会排除较短垂线。根据上面的分析，只要保留这条较短垂线，缩小宽度后形成的容器都不可能超过当前水量，因此被排除的组合不会遗漏更优解。

两个指针不断缩小搜索区间，并在排除每条边之前计算对应水量。所有可能成为最优解的组合都会被保留并检查，最终记录的最大值就是容器可以储存的最大水量。

## 复杂度分析

- 时间复杂度：`O(n)`。两个指针最多各移动 `n - 1` 次，整个数组只需扫描一次。
- 空间复杂度：`O(1)`。算法只使用两个指针和一个最大值变量。

## Go 代码

```go
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	result := 0

	for left < right {
		width := right - left
		currentArea := width * min(height[left], height[right])
		result = max(result, currentArea)

		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return result
}
```

## 小结

这道题的关键是同时观察容器的宽度和有效高度。双指针从最大宽度开始搜索，每次移动较短的一侧，安全排除不可能产生更大水量的组合。最终用 `O(n)` 时间和 `O(1)` 额外空间得到答案。
