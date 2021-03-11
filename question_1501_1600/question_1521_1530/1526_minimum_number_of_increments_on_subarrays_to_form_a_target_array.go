package question_1521_1530

// 1526. 形成目标数组的子数组最少增加次数
// https://leetcode-cn.com/problems/minimum-number-of-increments-on-subarrays-to-form-a-target-array/
// Topics: 线段树

func minNumberOperations(target []int) int {
	var res = target[0]
	for i := 1; i < len(target); i++ {
		if target[i] > target[i-1] {
			res += target[i] - target[i-1]
		}
	}
	return res
}
