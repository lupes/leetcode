package question_471_480

// 473. 火柴拼正方形
// https://leetcode-cn.com/problems/matchsticks-to-square/

func makesquare(nums []int) bool {
	r := int(0)
	for _, n := range nums {
		r += n
	}
	if r%4 != 0 {
		return false
	}
	a := r / 4
	return makeSquareHelp(nums, a, 0, 4)
}

func makeSquareHelp(nums []int, a, t, j int) bool {
	if t == 0 && j == 0 {
		return true
	}
	flag := false
	for i, n := range nums {
		if n != 0 {
			nums[i] = 0
			if t+n == a {
				flag = flag || makeSquareHelp(nums, a, 0, j-1)
			} else if t+n < a {
				flag = flag || makeSquareHelp(nums, a, t+n, j)
			}
			nums[i] = n
		}
	}
	return flag
}
