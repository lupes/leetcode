package question_1291_1300

// 1295. 统计位数为偶数的数字
// https://leetcode-cn.com/problems/find-numbers-with-even-number-of-digits/
// Topics: 数组

func findNumbers(nums []int) int {
	var res = 0
	for _, n := range nums {
		if (n >= 1000 && n <= 9999) || (n >= 10 && n <= 99) || n == 100000 {
			res += 1
		}
	}
	return res
}
