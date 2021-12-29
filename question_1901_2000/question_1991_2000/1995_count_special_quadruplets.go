package question_1991_2000

// 1995. 统计特殊四元组
// https://leetcode-cn.com/problems/count-special-quadruplets/
// Topic: 数组 枚举

func countQuadruplets(nums []int) int {
	var flag = make(map[int][]int)
	for i, n := range nums {
		flag[n] = append(flag[n], i)
	}

	var res, l = 0, len(nums)
	for a := 0; a < l-3; a++ {
		for b := a + 1; b < l-2; b++ {
			for c := b + 1; c < l-1; c++ {
				arr := flag[nums[a]+nums[b]+nums[c]]
				for _, i := range arr {
					if i > c {
						res++
					}
				}
			}
		}
	}
	return res
}
