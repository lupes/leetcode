package question_731_740

// 738. 单调递增的数字
// https://leetcode-cn.com/problems/monotone-increasing-digits
// Topics: 贪心算法

func monotoneIncreasingDigits(N int) int {
	var flag = make([]int, 0)
	for N > 0 {
		flag = append([]int{N % 10}, flag...)
		N /= 10
	}
	for i := 1; i < len(flag); i++ {
		if flag[i] < flag[i-1] {
			for j := i; j > 0; j-- {
				if flag[j] < flag[j-1] {
					flag[j] = 9
					flag[j-1] -= 1
				} else {
					break
				}
			}
			for j := i; j < len(flag); j++ {
				flag[j] = 9
			}
			break
		}
	}
	var res int
	for _, n := range flag {
		res = res*10 + n
	}
	return res
}
