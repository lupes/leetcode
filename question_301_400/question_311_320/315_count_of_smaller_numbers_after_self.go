package question_311_320

// 315. 计算右侧小于当前元素的个数
// https://leetcode-cn.com/problems/count-of-smaller-numbers-after-self
// Topics: 排序 树状数组 线段树 二分查找 分治算法

func countSmaller(nums []int) []int {
	var res, flag = make([]int, len(nums)), make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		flag = append(flag, nums[i])
		for k := len(flag) - 1; k > 0; k-- {
			if flag[k] <= flag[k-1] {
				flag[k], flag[k-1] = flag[k-1], flag[k]
			} else {
				res[i] = k
				break
			}
		}
	}
	return res
}
