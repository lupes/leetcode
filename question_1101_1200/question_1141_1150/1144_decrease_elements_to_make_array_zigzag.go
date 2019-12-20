package question_1141_1150

// 1144. 递减元素使数组呈锯齿状
// https://leetcode-cn.com/problems/decrease-elements-to-make-array-zigzag
// Topics: 数组

func movesToMakeZigzag(nums []int) int {
	var l, r1, r2 = len(nums), 0, 0
	for i, n := range nums {
		if i%2 == 0 {
			if i < l-1 && nums[i+1] <= n {
				r1 += n - nums[i+1] + 1
				n -= n - nums[i+1] + 1
			}
			if i > 0 && nums[i-1] <= n {
				r1 += n - nums[i-1] + 1
				n -= n - nums[i-1] + 1
			}
		} else {
			if i < l-1 && nums[i+1] <= n {
				r2 += n - nums[i+1] + 1
				n -= n - nums[i+1] + 1
			}
			if i > 0 && nums[i-1] <= n {
				r2 += n - nums[i-1] + 1
				n -= n - nums[i-1] + 1
			}
		}
	}
	if r1 > r2 {
		return r2
	}
	return r1
}
