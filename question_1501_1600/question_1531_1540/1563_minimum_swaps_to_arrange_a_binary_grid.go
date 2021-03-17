package question_1531_1540

// 1536. 排布二进制网格的最少交换次数
// https://leetcode-cn.com/problems/minimum-swaps-to-arrange-a-binary-grid/
// Topics: 贪心算法

func minSwaps(grid [][]int) int {
	size := len(grid)
	var flag = make(map[int]int, size)
	for i, row := range grid {
		flag[i+1] = 0
		for j := size - 1; j >= 0 && row[j] == 0; j-- {
			flag[i+1]++
		}
	}

	var res = 0
	for i := 1; i <= size; i++ {
		if flag[i] < size-i {
			f, find := flag[i], false
			for j := i + 1; j <= size; j++ {
				if flag[j] >= size-i {
					flag[i], flag[j], res, find = flag[j], f, res+j-i, true
					break
				} else {
					f, flag[j] = flag[j], f
				}
			}
			if !find {
				return -1
			}
		}
	}

	return res
}
