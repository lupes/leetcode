package question_801_810

// 807. 保持城市天际线
// https://leetcode-cn.com/problems/max-increase-to-keep-city-skyline
// Topics:

func maxIncreaseKeepingSkyline(grid [][]int) int {
	rowMax := make([]int, len(grid))
	colMax := make([]int, len(grid[0]))
	for i, r := range grid {
		for j, h := range r {
			if h > rowMax[i] {
				rowMax[i] = h
			}
			if h > colMax[j] {
				colMax[j] = h
			}
		}
	}
	var res = 0
	for i, r := range grid {
		for j, h := range r {
			min := rowMax[i]
			if min > colMax[j] {
				min = colMax[j]
			}
			res += min - h
		}
	}
	return res
}
