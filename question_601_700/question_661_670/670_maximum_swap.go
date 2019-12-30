package question_661_670

// 670. 最大交换
// https://leetcode-cn.com/problems/maximum-swap
// Topics: 数组 数学

func maximumSwap(num int) int {
	var tmp []int
	var res = 0
	for num > 0 {
		tmp = append([]int{num % 10}, tmp...)
		num /= 10
	}
	for i, n := range tmp {
		index, max := 0, 0
		for j, m := range tmp[i+1:] {
			if m > n && m >= max {
				max, index = m, j+i+1
			}
		}
		if max != 0 {
			tmp[i], tmp[index] = tmp[index], tmp[i]
			break
		}
	}
	for _, n := range tmp {
		res = res*10 + n
	}
	return res
}
