package question_1551_1560

// 1560. 圆形赛道上经过次数最多的扇区
// https://leetcode-cn.com/problems/most-visited-sector-in-a-circular-track/
// Topics: 数组

func mostVisited(n int, rounds []int) []int {
	var flag = make([]int, n+1)
	flag[rounds[0]%n]++
	for i := 1; i < len(rounds); i++ {
		a, b := rounds[i-1], rounds[i]
		if b < a {
			b += n
		}
		for j := a + 1; j <= b; j++ {
			flag[j%n]++
		}
	}
	flag[n] = flag[0]
	var max, res = 1, make([]int, 0)
	for i := 1; i <= n; i++ {
		if flag[i] > max {
			max = flag[i]
			res = []int{i}
		} else if flag[i] == max {
			res = append(res, i)
		}
	}
	return res
}
