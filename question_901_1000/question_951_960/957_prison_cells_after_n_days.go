package question_951_960

// 957. N 天后的牢房
// https://leetcode-cn.com/problems/prison-cells-after-n-days
// Topics: 哈希表

func prisonAfterNDays(cells []int, N int) []int {
	N = helper(cells, N)
	for i := 0; i < N; i++ {
		cells = next(cells)
	}
	return cells
}

func helper(cells []int, N int) int {
	var flag = make(map[int]int, 0)
	for i := 0; i < N; i++ {
		cells = next(cells)
		f := cellsToInt(cells)
		if _, ok := flag[f]; ok {
			N = N % i
			if N == 0 {
				N = i
			}
			return N
		} else {
			flag[f] = i
		}
	}
	return N
}

func next(cells []int) []int {
	var next = make([]int, 8)
	for i := 1; i < 7; i++ {
		if cells[i-1] == cells[i+1] {
			next[i] = 1
		}
	}
	return next
}

func cellsToInt(cells []int) int {
	var res = 0
	for i := 1; i < 7; i++ {
		res = (res << 1) + cells[i]
	}
	return res
}
