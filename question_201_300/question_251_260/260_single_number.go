package question_251_260

// 260. 只出现一次的数字 III
// https://leetcode-cn.com/problems/single-number-iii/

func singleNumber(nums []int) []int {
	t, i := 0, uint(0)
	for _, n := range nums {
		t ^= n
	}
	for i = 0; (t>>i)&1 != 1; i++ {
	}

	var a, b = 0, 0
	for _, n := range nums {
		if (n>>i)&1 == 0 {
			a ^= n
		} else {
			b ^= n
		}
	}

	return []int{a, b}
}
