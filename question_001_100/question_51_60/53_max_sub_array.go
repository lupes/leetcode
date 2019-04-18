package question_51_60

func maxSubArray(nums []int) int {
	res := nums[0]
	for i, n := range nums {
		max := n
		t := n
		for _, m := range nums[i+1:] {
			t = t + m
			if t > max {
				max = t
			}
		}
		if max > res {
			res = max
		}
	}
	return res
}
