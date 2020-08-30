package question_1151_1160

// 1155. 掷骰子的N种方法
// https://leetcode-cn.com/problems/number-of-dice-rolls-with-target-sum
// Topics: 动态规划

func numRollsToTarget(d int, f int, target int) int {
	var flag = make([][]int, d+1)
	for i := range flag {
		flag[i] = make([]int, target+1)
	}
	flag[0][0] = 1
	for i := 1; i <= d; i++ {
		for j := 1; j <= f; j++ {
			for k := range flag[i] {
				if k-j >= 0 {
					flag[i][k] = (flag[i][k] + flag[i-1][k-j]) % (1e9 + 7)
				}
			}
		}
	}
	return flag[d][target]
}
