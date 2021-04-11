package question_1611_1620

// 1615. 最大网络秩
// https://leetcode-cn.com/problems/maximal-network-rank/
// Topics: 图

func maximalNetworkRank(n int, roads [][]int) int {
	var cities = make([]map[int]struct{}, n)
	for i := 0; i < n; i++ {
		cities[i] = make(map[int]struct{})
	}
	for _, road := range roads {
		cities[road[0]][road[1]] = struct{}{}
		cities[road[1]][road[0]] = struct{}{}
	}

	var max = 0
	for i := range cities {
		for j := i + 1; j < n; j++ {
			l := len(cities[i]) + len(cities[j])
			if _, ok := cities[i][j]; ok {
				l -= 1
			}
			if l > max {
				max = l
			}
		}
	}
	return max
}
