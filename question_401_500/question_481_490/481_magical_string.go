package question_481_490

// 481. 神奇字符串
// https://leetcode-cn.com/problems/magical-string
// Topics:

func magicalString(n int) int {
	var res, flag = 1, make([]byte, 0, n+1)
	flag = append(flag, 1, 2, 2)
	for low, fast := 2, 3; fast < n; low++ {
		for i := 0; i < int(flag[low]); i++ {
			flag = append(flag, flag[fast-1]^3)
		}
		if flag[fast-1] == 2 {
			res += int(flag[low])
		}
		fast += int(flag[low])
	}
	if len(flag) > n && flag[n] == 1 {
		res -= 1
	}
	return res
}
