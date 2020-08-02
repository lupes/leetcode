package question_731_740

// 738. 单调递增的数字
// https://leetcode-cn.com/problems/monotone-increasing-digits
// Topics: 贪心算法

func monotoneIncreasingDigits(N int) int {
	var flag = make([]int, 0)
	for N > 0 {
		flag = append(flag, N%10)
		N /= 10
	}
	for i := 1; i < len(flag); i++ {
		if flag[i] > flag[i-1] {
			flag[i] -= 1
			for j := i - 1; j >= 0 && flag[j] != 9; j-- {
				flag[j] = 9
			}
		}
	}
	var res int
	for i := len(flag) - 1; i >= 0; i-- {
		res = res*10 + flag[i]
	}
	return res
}
