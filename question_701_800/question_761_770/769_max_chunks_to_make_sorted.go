package question_761_770

// 769. 最多能完成排序的块
// https://leetcode-cn.com/problems/max-chunks-to-make-sorted
// Topics: 数组

func maxChunksToSorted(arr []int) int {
	var flag = make([]int, len(arr))
	for i, n := range arr {
		flag[n] = i
	}

	var res = 0
	for start, end := 0, 0; start < len(arr); start, res = end+1, res+1 {
		end = arr[start]
		for i := start; i <= end; i++ {
			if arr[i] > end {
				end = arr[i]
			}
		}
	}
	return res
}
