# Go 刷题练习目录

沿用此前整理的 Hot 100 其余 **94 题**，每题提供空白函数、参数和返回值约定、可运行测试。已在博客整理的 3、11、15、42、128、283 题未重复加入本批目录。

## 快速开始

要求 Go 1.21 或更高版本，使用 64 位环境。只使用 Go 标准库，无需安装第三方依赖。

在项目根目录执行：

```sh
cd leetcode
# 先写 hot100/p0001_two_sum/solution.go，然后运行单题：
go test ./hot100/p0001_two_sum -v -count=1 -timeout=30s
# 运行全部题目：
go test ./... -v -count=1 -timeout=60s
# 只检查所有包能否编译：
go test ./... -run '^$'
```

也可以进入某一题目录，运行 `go test -v`。VS Code、GoLand 等编辑器可在测试函数旁直接运行。

## 文件约定

```text
leetcode/
  go.mod
  README.md
  internal/testutil/        # 将模板 TODO 标记为 SKIP，其他错误继续失败
  hot100/
    p0001_two_sum/
      README.md            # 题意、示例、函数签名、参数、返回值、运行命令
      solution.go          # 你在这里写代码
      solution_test.go     # 样例、边界及部分随机/穷举测试
    p0002_add_two_numbers/
      types.go             # 链表、树节点或设计题结构体
      ...
```

- 每题是独立的 `solution` 包，函数名与 LeetCode 保持一致。第 48 和 189 题都能使用 `rotate`，各设计题都能使用 `Constructor`。
- 普通题保留函数签名并编写函数体；设计题还需要在 `types.go` 中添加实例字段。辅助函数可写在同一题目录。
- 想练多种解法，可以将它们写成不同辅助函数，由固定入口函数选择调用；同一套测试可重复使用。
- 未实现的方法包含 `panic("TODO")`。测试遇到它会显示 `SKIP`；填好实现并移除标记后，测试自动校验结果。`SKIP` 表示待完成，不能算作解题通过。全量命令退出码为 0 也可能包含跳过项，请使用 `-v` 查看明细。
- 越界、空指针及其他 panic 都会失败。测试不会因为一般性错误而自动跳过。
- 本地测试包含题目示例、边界及部分穷举/固定种子随机输入。代码正确性需结合约束判断，测试通过不证明满足题目的时间或空间要求。
- 链表、树的测试会比较节点身份或结构；无需自己解析 LeetCode 的文本输入。无序答案按题意比较，允许合理的返回顺序。
- 一些题额外覆盖空输入或恢复原结构的边界，相关约定写在各题 README；原题完整条件以题目链接为准。
- 提交到 LeetCode 时复制自己的实现和必要的辅助代码。平台已有节点类型时保留一份定义；本地测试代码无需提交。

## 题目目录

### 哈希

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 1 | [两数之和](hot100/p0001_two_sum/README.md) | [solution.go](hot100/p0001_two_sum/solution.go) |
| 49 | [字母异位词分组](hot100/p0049_group_anagrams/README.md) | [solution.go](hot100/p0049_group_anagrams/solution.go) |

### 滑动窗口

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 438 | [找到字符串中所有字母异位词](hot100/p0438_find_all_anagrams_in_a_string/README.md) | [solution.go](hot100/p0438_find_all_anagrams_in_a_string/solution.go) |

### 子串

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 560 | [和为 K 的子数组](hot100/p0560_subarray_sum_equals_k/README.md) | [solution.go](hot100/p0560_subarray_sum_equals_k/solution.go) |
| 239 | [滑动窗口最大值](hot100/p0239_sliding_window_maximum/README.md) | [solution.go](hot100/p0239_sliding_window_maximum/solution.go) |
| 76 | [最小覆盖子串](hot100/p0076_minimum_window_substring/README.md) | [solution.go](hot100/p0076_minimum_window_substring/solution.go) |

### 普通数组

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 53 | [最大子数组和](hot100/p0053_maximum_subarray/README.md) | [solution.go](hot100/p0053_maximum_subarray/solution.go) |
| 56 | [合并区间](hot100/p0056_merge_intervals/README.md) | [solution.go](hot100/p0056_merge_intervals/solution.go) |
| 189 | [轮转数组](hot100/p0189_rotate_array/README.md) | [solution.go](hot100/p0189_rotate_array/solution.go) |
| 238 | [除了自身以外数组的乘积](hot100/p0238_product_of_array_except_self/README.md) | [solution.go](hot100/p0238_product_of_array_except_self/solution.go) |
| 41 | [缺失的第一个正数](hot100/p0041_first_missing_positive/README.md) | [solution.go](hot100/p0041_first_missing_positive/solution.go) |

### 矩阵

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 73 | [矩阵置零](hot100/p0073_set_matrix_zeroes/README.md) | [solution.go](hot100/p0073_set_matrix_zeroes/solution.go) |
| 54 | [螺旋矩阵](hot100/p0054_spiral_matrix/README.md) | [solution.go](hot100/p0054_spiral_matrix/solution.go) |
| 48 | [旋转图像](hot100/p0048_rotate_image/README.md) | [solution.go](hot100/p0048_rotate_image/solution.go) |
| 240 | [搜索二维矩阵 II](hot100/p0240_search_a_2d_matrix_ii/README.md) | [solution.go](hot100/p0240_search_a_2d_matrix_ii/solution.go) |

### 链表

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 160 | [相交链表](hot100/p0160_intersection_of_two_linked_lists/README.md) | [solution.go](hot100/p0160_intersection_of_two_linked_lists/solution.go) |
| 206 | [反转链表](hot100/p0206_reverse_linked_list/README.md) | [solution.go](hot100/p0206_reverse_linked_list/solution.go) |
| 234 | [回文链表](hot100/p0234_palindrome_linked_list/README.md) | [solution.go](hot100/p0234_palindrome_linked_list/solution.go) |
| 141 | [环形链表](hot100/p0141_linked_list_cycle/README.md) | [solution.go](hot100/p0141_linked_list_cycle/solution.go) |
| 142 | [环形链表 II](hot100/p0142_linked_list_cycle_ii/README.md) | [solution.go](hot100/p0142_linked_list_cycle_ii/solution.go) |
| 21 | [合并两个有序链表](hot100/p0021_merge_two_sorted_lists/README.md) | [solution.go](hot100/p0021_merge_two_sorted_lists/solution.go) |
| 2 | [两数相加](hot100/p0002_add_two_numbers/README.md) | [solution.go](hot100/p0002_add_two_numbers/solution.go) |
| 19 | [删除链表的倒数第 N 个结点](hot100/p0019_remove_nth_node_from_end_of_list/README.md) | [solution.go](hot100/p0019_remove_nth_node_from_end_of_list/solution.go) |
| 24 | [两两交换链表中的节点](hot100/p0024_swap_nodes_in_pairs/README.md) | [solution.go](hot100/p0024_swap_nodes_in_pairs/solution.go) |
| 25 | [K 个一组翻转链表](hot100/p0025_reverse_nodes_in_k_group/README.md) | [solution.go](hot100/p0025_reverse_nodes_in_k_group/solution.go) |
| 138 | [随机链表的复制](hot100/p0138_copy_list_with_random_pointer/README.md) | [solution.go](hot100/p0138_copy_list_with_random_pointer/solution.go) |
| 148 | [排序链表](hot100/p0148_sort_list/README.md) | [solution.go](hot100/p0148_sort_list/solution.go) |
| 23 | [合并 K 个升序链表](hot100/p0023_merge_k_sorted_lists/README.md) | [solution.go](hot100/p0023_merge_k_sorted_lists/solution.go) |
| 146 | [LRU 缓存](hot100/p0146_lru_cache/README.md) | [solution.go](hot100/p0146_lru_cache/solution.go) |

### 二叉树

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 94 | [二叉树的中序遍历](hot100/p0094_binary_tree_inorder_traversal/README.md) | [solution.go](hot100/p0094_binary_tree_inorder_traversal/solution.go) |
| 104 | [二叉树的最大深度](hot100/p0104_maximum_depth_of_binary_tree/README.md) | [solution.go](hot100/p0104_maximum_depth_of_binary_tree/solution.go) |
| 226 | [翻转二叉树](hot100/p0226_invert_binary_tree/README.md) | [solution.go](hot100/p0226_invert_binary_tree/solution.go) |
| 101 | [对称二叉树](hot100/p0101_symmetric_tree/README.md) | [solution.go](hot100/p0101_symmetric_tree/solution.go) |
| 543 | [二叉树的直径](hot100/p0543_diameter_of_binary_tree/README.md) | [solution.go](hot100/p0543_diameter_of_binary_tree/solution.go) |
| 102 | [二叉树的层序遍历](hot100/p0102_binary_tree_level_order_traversal/README.md) | [solution.go](hot100/p0102_binary_tree_level_order_traversal/solution.go) |
| 108 | [将有序数组转换为二叉搜索树](hot100/p0108_convert_sorted_array_to_binary_search_tree/README.md) | [solution.go](hot100/p0108_convert_sorted_array_to_binary_search_tree/solution.go) |
| 98 | [验证二叉搜索树](hot100/p0098_validate_binary_search_tree/README.md) | [solution.go](hot100/p0098_validate_binary_search_tree/solution.go) |
| 230 | [二叉搜索树中第 K 小的元素](hot100/p0230_kth_smallest_element_in_a_bst/README.md) | [solution.go](hot100/p0230_kth_smallest_element_in_a_bst/solution.go) |
| 199 | [二叉树的右视图](hot100/p0199_binary_tree_right_side_view/README.md) | [solution.go](hot100/p0199_binary_tree_right_side_view/solution.go) |
| 114 | [二叉树展开为链表](hot100/p0114_flatten_binary_tree_to_linked_list/README.md) | [solution.go](hot100/p0114_flatten_binary_tree_to_linked_list/solution.go) |
| 105 | [从前序与中序遍历序列构造二叉树](hot100/p0105_construct_binary_tree_from_preorder_and_inorder_traversal/README.md) | [solution.go](hot100/p0105_construct_binary_tree_from_preorder_and_inorder_traversal/solution.go) |
| 437 | [路径总和 III](hot100/p0437_path_sum_iii/README.md) | [solution.go](hot100/p0437_path_sum_iii/solution.go) |
| 236 | [二叉树的最近公共祖先](hot100/p0236_lowest_common_ancestor_of_a_binary_tree/README.md) | [solution.go](hot100/p0236_lowest_common_ancestor_of_a_binary_tree/solution.go) |
| 124 | [二叉树中的最大路径和](hot100/p0124_binary_tree_maximum_path_sum/README.md) | [solution.go](hot100/p0124_binary_tree_maximum_path_sum/solution.go) |

### 图论

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 200 | [岛屿数量](hot100/p0200_number_of_islands/README.md) | [solution.go](hot100/p0200_number_of_islands/solution.go) |
| 994 | [腐烂的橘子](hot100/p0994_rotting_oranges/README.md) | [solution.go](hot100/p0994_rotting_oranges/solution.go) |
| 207 | [课程表](hot100/p0207_course_schedule/README.md) | [solution.go](hot100/p0207_course_schedule/solution.go) |
| 208 | [实现 Trie (前缀树)](hot100/p0208_implement_trie_prefix_tree/README.md) | [solution.go](hot100/p0208_implement_trie_prefix_tree/solution.go) |

### 回溯

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 46 | [全排列](hot100/p0046_permutations/README.md) | [solution.go](hot100/p0046_permutations/solution.go) |
| 78 | [子集](hot100/p0078_subsets/README.md) | [solution.go](hot100/p0078_subsets/solution.go) |
| 17 | [电话号码的字母组合](hot100/p0017_letter_combinations_of_a_phone_number/README.md) | [solution.go](hot100/p0017_letter_combinations_of_a_phone_number/solution.go) |
| 39 | [组合总和](hot100/p0039_combination_sum/README.md) | [solution.go](hot100/p0039_combination_sum/solution.go) |
| 22 | [括号生成](hot100/p0022_generate_parentheses/README.md) | [solution.go](hot100/p0022_generate_parentheses/solution.go) |
| 79 | [单词搜索](hot100/p0079_word_search/README.md) | [solution.go](hot100/p0079_word_search/solution.go) |
| 131 | [分割回文串](hot100/p0131_palindrome_partitioning/README.md) | [solution.go](hot100/p0131_palindrome_partitioning/solution.go) |
| 51 | [N 皇后](hot100/p0051_n_queens/README.md) | [solution.go](hot100/p0051_n_queens/solution.go) |

### 二分查找

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 35 | [搜索插入位置](hot100/p0035_search_insert_position/README.md) | [solution.go](hot100/p0035_search_insert_position/solution.go) |
| 74 | [搜索二维矩阵](hot100/p0074_search_a_2d_matrix/README.md) | [solution.go](hot100/p0074_search_a_2d_matrix/solution.go) |
| 34 | [在排序数组中查找元素的第一个和最后一个位置](hot100/p0034_find_first_and_last_position_of_element_in_sorted_array/README.md) | [solution.go](hot100/p0034_find_first_and_last_position_of_element_in_sorted_array/solution.go) |
| 33 | [搜索旋转排序数组](hot100/p0033_search_in_rotated_sorted_array/README.md) | [solution.go](hot100/p0033_search_in_rotated_sorted_array/solution.go) |
| 153 | [寻找旋转排序数组中的最小值](hot100/p0153_find_minimum_in_rotated_sorted_array/README.md) | [solution.go](hot100/p0153_find_minimum_in_rotated_sorted_array/solution.go) |
| 4 | [寻找两个正序数组的中位数](hot100/p0004_median_of_two_sorted_arrays/README.md) | [solution.go](hot100/p0004_median_of_two_sorted_arrays/solution.go) |

### 栈

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 20 | [有效的括号](hot100/p0020_valid_parentheses/README.md) | [solution.go](hot100/p0020_valid_parentheses/solution.go) |
| 155 | [最小栈](hot100/p0155_min_stack/README.md) | [solution.go](hot100/p0155_min_stack/solution.go) |
| 394 | [字符串解码](hot100/p0394_decode_string/README.md) | [solution.go](hot100/p0394_decode_string/solution.go) |
| 739 | [每日温度](hot100/p0739_daily_temperatures/README.md) | [solution.go](hot100/p0739_daily_temperatures/solution.go) |
| 84 | [柱状图中最大的矩形](hot100/p0084_largest_rectangle_in_histogram/README.md) | [solution.go](hot100/p0084_largest_rectangle_in_histogram/solution.go) |

### 堆

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 215 | [数组中的第K个最大元素](hot100/p0215_kth_largest_element_in_an_array/README.md) | [solution.go](hot100/p0215_kth_largest_element_in_an_array/solution.go) |
| 347 | [前 K 个高频元素](hot100/p0347_top_k_frequent_elements/README.md) | [solution.go](hot100/p0347_top_k_frequent_elements/solution.go) |
| 295 | [数据流的中位数](hot100/p0295_find_median_from_data_stream/README.md) | [solution.go](hot100/p0295_find_median_from_data_stream/solution.go) |

### 贪心算法

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 121 | [买卖股票的最佳时机](hot100/p0121_best_time_to_buy_and_sell_stock/README.md) | [solution.go](hot100/p0121_best_time_to_buy_and_sell_stock/solution.go) |
| 55 | [跳跃游戏](hot100/p0055_jump_game/README.md) | [solution.go](hot100/p0055_jump_game/solution.go) |
| 45 | [跳跃游戏 II](hot100/p0045_jump_game_ii/README.md) | [solution.go](hot100/p0045_jump_game_ii/solution.go) |
| 763 | [划分字母区间](hot100/p0763_partition_labels/README.md) | [solution.go](hot100/p0763_partition_labels/solution.go) |

### 动态规划

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 70 | [爬楼梯](hot100/p0070_climbing_stairs/README.md) | [solution.go](hot100/p0070_climbing_stairs/solution.go) |
| 118 | [杨辉三角](hot100/p0118_pascals_triangle/README.md) | [solution.go](hot100/p0118_pascals_triangle/solution.go) |
| 198 | [打家劫舍](hot100/p0198_house_robber/README.md) | [solution.go](hot100/p0198_house_robber/solution.go) |
| 279 | [完全平方数](hot100/p0279_perfect_squares/README.md) | [solution.go](hot100/p0279_perfect_squares/solution.go) |
| 322 | [零钱兑换](hot100/p0322_coin_change/README.md) | [solution.go](hot100/p0322_coin_change/solution.go) |
| 139 | [单词拆分](hot100/p0139_word_break/README.md) | [solution.go](hot100/p0139_word_break/solution.go) |
| 300 | [最长递增子序列](hot100/p0300_longest_increasing_subsequence/README.md) | [solution.go](hot100/p0300_longest_increasing_subsequence/solution.go) |
| 152 | [乘积最大子数组](hot100/p0152_maximum_product_subarray/README.md) | [solution.go](hot100/p0152_maximum_product_subarray/solution.go) |
| 416 | [分割等和子集](hot100/p0416_partition_equal_subset_sum/README.md) | [solution.go](hot100/p0416_partition_equal_subset_sum/solution.go) |
| 32 | [最长有效括号](hot100/p0032_longest_valid_parentheses/README.md) | [solution.go](hot100/p0032_longest_valid_parentheses/solution.go) |

### 多维动态规划

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 62 | [不同路径](hot100/p0062_unique_paths/README.md) | [solution.go](hot100/p0062_unique_paths/solution.go) |
| 64 | [最小路径和](hot100/p0064_minimum_path_sum/README.md) | [solution.go](hot100/p0064_minimum_path_sum/solution.go) |
| 5 | [最长回文子串](hot100/p0005_longest_palindromic_substring/README.md) | [solution.go](hot100/p0005_longest_palindromic_substring/solution.go) |
| 1143 | [最长公共子序列](hot100/p1143_longest_common_subsequence/README.md) | [solution.go](hot100/p1143_longest_common_subsequence/solution.go) |
| 72 | [编辑距离](hot100/p0072_edit_distance/README.md) | [solution.go](hot100/p0072_edit_distance/solution.go) |

### 技巧

| 题号 | 题目 | 入口文件 |
|---:|---|---|
| 136 | [只出现一次的数字](hot100/p0136_single_number/README.md) | [solution.go](hot100/p0136_single_number/solution.go) |
| 169 | [多数元素](hot100/p0169_majority_element/README.md) | [solution.go](hot100/p0169_majority_element/solution.go) |
| 75 | [颜色分类](hot100/p0075_sort_colors/README.md) | [solution.go](hot100/p0075_sort_colors/solution.go) |
| 31 | [下一个排列](hot100/p0031_next_permutation/README.md) | [solution.go](hot100/p0031_next_permutation/solution.go) |
| 287 | [寻找重复数](hot100/p0287_find_the_duplicate_number/README.md) | [solution.go](hot100/p0287_find_the_duplicate_number/solution.go) |
