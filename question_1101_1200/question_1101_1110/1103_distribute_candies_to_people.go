package question_1101_1110

// 1103. 分糖果 II
// https://leetcode-cn.com/problems/distribute-candies-to-people
// Topics: 数学

func distributeCandies(candies int, num_people int) []int {
	var res = make([]int, num_people)
	for i, num := 0, 1; candies > 0; i, num = i+1, num+1 {
		if candies-num < 0 {
			num = candies
		}
		res[i%num_people], candies = res[i%num_people]+num, candies-num
	}
	return res
}
