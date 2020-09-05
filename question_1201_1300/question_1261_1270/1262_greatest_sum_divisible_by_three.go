package question_1261_1270

// 1262. 可被三整除的最大和
// https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/
// Topics: 动态规划

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxSumDivThree(nums []int) int {
	var v0, v1, v2 int
	for _, n := range nums {
		switch n % 3 {
		case 0:
			v0, v1, v2 = v0+n, v1+n, v2+n
		case 1:
			t := v0
			if (n+v2)%3 == 0 {
				v0 = max(v0, n+v2)
			}
			if (n+v1)%3 == 2 {
				v2 = max(v2, v1+n)
			}
			if (n+t)%3 == 1 {
				v1 = max(v1, t+n)
			}
		case 2:
			t := v0
			if (n+v1)%3 == 0 {
				v0 = max(v0, v1+n)
			}
			if (n+v2)%3 == 1 {
				v1 = max(v1, v2+n)
			}
			if (n+t)%3 == 2 {
				v2 = max(v2, t+n)
			}
		}
	}
	return v0
}
