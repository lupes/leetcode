package question_211_220

// 218. 天际线问题
// https://leetcode-cn.com/problems/the-skyline-problem

func getSkyline(buildings [][]int) [][]int {
	if len(buildings) == 0 {
		return [][]int{}
	}
	var lines [][]int
	for _, building := range buildings {
		t := true
		for i, line := range lines {
			if line[0] <= building[0] && building[0] <= line[1] {
				if building[1] > line[1] {
					lines[i][1] = building[1]
				}
				t = false
			}

			if line[0] == building[0] && building[2] > line[2] {
				lines[i][2] = building[2]
			}
		}
		if t {
			lines = append(lines, []int{building[0], building[1], building[2]})
		}
	}
	var res [][]int
	for _, line := range lines {
		res = append(res, []int{line[0], line[2]})
		height := line[2]
		for i := line[0]; i <= line[1]; i++ {
			x, y := i, 0
			for _, building := range buildings {
				if building[0] > x {
					break
				}
				if building[0] <= x && x <= building[1] && building[2] > y {
					y = building[2]
				}
			}
			if y > height {
				height = y
				res = append(res, []int{x, y})
			} else if y < height {
				height = y
				res = append(res, []int{x - 1, y})
			}
		}
		res = append(res, []int{line[1], 0})
	}
	return res
}
