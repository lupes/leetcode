package question_1401_1410

// 1405. 最长快乐字符串
// https://leetcode-cn.com/problems/longest-happy-string/
// Topics: 贪心算法 动态规划

func longestDiverseString(a int, b int, c int) string {
	var res = make([]byte, 2, 2+a+b+c+1)
	i := 2
	for ; a > 0 || b > 0 || c > 0; i++ {
		if res[i-1] == 'a' && res[i-2] == 'a' && b >= c && b > 0 {
			res = append(res, 'b')
			b--
		} else if res[i-1] == 'a' && res[i-2] == 'a' && c > 0 {
			res = append(res, 'c')
			c--
		} else if res[i-1] == 'b' && res[i-2] == 'b' && a >= c && a > 0 {
			res = append(res, 'a')
			a--
		} else if res[i-1] == 'b' && res[i-2] == 'b' && c > 0 {
			res = append(res, 'c')
			c--
		} else if res[i-1] == 'c' && res[i-2] == 'c' && a >= b && a > 0 {
			res = append(res, 'a')
			a--
		} else if res[i-1] == 'c' && res[i-2] == 'c' && b > 0 {
			res = append(res, 'b')
			b--
		} else if (res[i-1] != 'a' || res[i-2] != 'a') && a >= b && a >= c && a > 0 {
			res = append(res, 'a')
			a--
		} else if (res[i-1] != 'b' || res[i-2] != 'b') && b >= a && b >= c && b > 0 {
			res = append(res, 'b')
			b--
		} else if (res[i-1] != 'c' || res[i-2] != 'c') && c > 0 {
			res = append(res, 'c')
			c--
		} else {
			return string(res[2:i])
		}
	}
	return string(res[2:i])
}
