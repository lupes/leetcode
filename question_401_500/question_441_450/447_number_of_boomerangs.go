package question_441_450

// 447. 回旋镖的数量
// https://leetcode-cn.com/problems/number-of-boomerangs
// Topics: 哈希表

func numberOfBoomerangs(points [][]int) int {
	var res = 0
	for i := 0; i < len(points); i++ {
		var flag = make(map[int]int, 0)
		for j := 0; j < i; j++ {
			a := points[j][0] - points[i][0]
			b := points[j][1] - points[i][1]
			flag[a*a+b*b]++
		}
		for j := i + 1; j < len(points); j++ {
			a := points[j][0] - points[i][0]
			b := points[j][1] - points[i][1]
			flag[a*a+b*b]++
		}
		for _, n := range flag {
			if n > 1 {
				res += (n - 1) * n / 2 * 2
			}
		}
	}
	return res
}
