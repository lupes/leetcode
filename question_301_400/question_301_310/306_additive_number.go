package question_301_310

// 306. 累加数
// https://leetcode-cn.com/problems/additive-number
// Topics: 回溯算法

func isAdditiveNumber(num string) bool {
	var nums = make([]int, len(num))
	for i, c := range num {
		nums[i] = int(c - '0')
	}

	for i := 1; i < len(num)/2+1; i++ {
		for j := 1; j < len(num)/2+1; j++ {
			if isAdditiveNumberHelper(nums, 0, i, i+j) {
				return true
			}
		}
	}
	return false
}

func isAdditiveNumberHelper(nums []int, s, i, j int) bool {
	if j == len(nums) && s != 0 {
		return true
	} else if j == len(nums) {
		return false
	}
	if j-i > 1 && nums[i] == 0 {
		return false
	}
	t := i - s
	if j-i > i-s {
		t = j - i
	}
	if !(t > 1 && nums[j] == 0) && j+t <= len(nums) && equal(nums[s:i], nums[i:j], nums[j:j+t]) {
		return isAdditiveNumberHelper(nums, i, j, j+t)
	} else if !(t > 0 && nums[j] == 0) && j+t+1 <= len(nums) && equal(nums[s:i], nums[i:j], nums[j:j+t+1]) {
		return isAdditiveNumberHelper(nums, i, j, j+t+1)
	}
	return false
}

func equal(a, b, c []int) bool {
	t := 0
	for i := 0; i < len(a) || i < len(b) || i < len(c); i++ {
		var an, bn, cn int
		if i < len(a) {
			an = a[len(a)-i-1]
		}
		if i < len(b) {
			bn = b[len(b)-i-1]
		}
		if i < len(c) {
			cn = c[len(c)-i-1]
		}
		sum := an + bn + t
		if sum > 9 {
			t = 1
			sum -= 10
		} else {
			t = 0
		}
		if cn != sum {
			return false
		}
	}
	return t == 0
}
