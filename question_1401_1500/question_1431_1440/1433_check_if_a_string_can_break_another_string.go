package question_1431_1440

// 1433. 检查一个字符串是否可以打破另一个字符串
// https://leetcode-cn.com/problems/check-if-a-string-can-break-another-string/
// Topics: 贪心算法 字符串

func checkIfCanBreak(s1 string, s2 string) bool {
	var s1Flag, s2Flag [26]int
	for i, n := range s1 {
		s1Flag[n-'a']++
		s2Flag[s2[i]-'a']++
	}

	sum1, sum2, flag1, flag2 := 0, 0, true, true
	for i := 25; i >= 0; i-- {
		sum1 += s1Flag[i]
		sum2 += s2Flag[i]
		if (sum1 > sum2 && !flag1) || (sum2 > sum1 && !flag2) {
			return false
		} else if sum1 == sum2 && (flag1 || flag2) {
			sum1, sum2 = 0, 0
		} else if sum1 > sum2 && flag1 {
			sum1 -= sum2
			sum2 = 0
			flag2 = false
		} else if sum2 > sum1 && flag2 {
			sum2 -= sum1
			sum1 = 0
			flag1 = false
		}
	}

	return true
}

func checkIfCanBreakHelper(s string, arr [26]int) bool {
	for _, n := range s {
		flag := false
		for j := n - 'a'; j < 26; j++ {
			if arr[j] > 0 {
				arr[j]--
				flag = true
				break
			}
		}
		if !flag {
			return false
		}
	}
	return true
}
