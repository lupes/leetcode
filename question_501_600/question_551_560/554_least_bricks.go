package question_551_560

// 554. 砖墙
// https://leetcode-cn.com/problems/brick-wall/

func leastBricks(wall [][]int) int {
	var tmp, width, max, height = make(map[int]int), 0, 0, len(wall)
	for _, n := range wall {
		width = 0
		for i := 0; i < len(n)-1; i++ {
			width += n[i]
			tmp[width]++
		}
	}
	for _, v := range tmp {
		if v > max {
			max = v
		}
	}
	return height - max
}
