package question_401_410

// 407. 接雨水 II
// https://leetcode-cn.com/problems/trapping-rain-water-ii
// Topics: 堆 广度优先搜索

var dirs = []int{-1, 0, 1, 0, -1}

func trapRainWater(heightMap [][]int) int {
	m, n := len(heightMap), len(heightMap[0])
	maxHeight := 0
	for _, row := range heightMap {
		for _, h := range row {
			maxHeight = max(maxHeight, h)
		}
	}

	water := make([][]int, m)
	for i := range water {
		water[i] = make([]int, n)
		for j := range water[i] {
			water[i][j] = maxHeight
		}
	}
	type pair struct{ x, y int }
	var q []pair
	for i, row := range heightMap {
		for j, h := range row {
			if (i == 0 || i == m-1 || j == 0 || j == n-1) && h < water[i][j] {
				water[i][j] = h
				q = append(q, pair{i, j})
			}
		}
	}

	for len(q) > 0 {
		p := q[0]
		q = q[1:]
		x, y := p.x, p.y
		for i := 0; i < 4; i++ {
			nx, ny := x+dirs[i], y+dirs[i+1]
			if 0 <= nx && nx < m && 0 <= ny && ny < n && water[nx][ny] > water[x][y] && water[nx][ny] > heightMap[nx][ny] {
				water[nx][ny] = max(water[x][y], heightMap[nx][ny])
				q = append(q, pair{nx, ny})
			}
		}
	}

	var res int
	for i, row := range heightMap {
		for j, h := range row {
			res += water[i][j] - h
		}
	}
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
