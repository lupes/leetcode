package question_1401_1410

// 1405. 最长快乐字符串
// https://leetcode-cn.com/problems/longest-happy-string/
// Topics: 贪心算法 动态规划

func helper(flag1, flag2, x, y, z byte, xs, ys, zs int) bool {
	return xs > 0 && ((flag1 == y && flag2 == y && xs >= zs) ||
		(flag1 == z && flag2 == z && xs >= ys && xs > 0) ||
		((flag1 != x || flag2 != x) && xs >= ys && xs >= zs))
}

func longestDiverseString(a int, b int, c int) string {
	var res = make([]byte, 2, 2+a+b+c+1)
	i := 2
	for ; a > 0 || b > 0 || c > 0; i++ {
		if helper(res[i-1], res[i-2], 'a', 'b', 'c', a, b, c) {
			res = append(res, 'a')
			a--
		} else if helper(res[i-1], res[i-2], 'b', 'a', 'c', b, a, c) {
			res = append(res, 'b')
			b--
		} else if helper(res[i-1], res[i-2], 'c', 'b', 'a', c, b, a) {
			res = append(res, 'c')
			c--
		} else {
			return string(res[2:i])
		}
	}
	return string(res[2:i])
}
