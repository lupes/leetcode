package question_121_130

// 128. 最长连续序列
// https://leetcode-cn.com/problems/longest-consecutive-sequence
// Topics: 并查集 数组

func longestConsecutive(nums []int) int {
	var res, flag = 0, make(map[int]int)
	for _, n := range nums {
		if _, ok := flag[n]; !ok {
			left, right := flag[n-1], flag[n+1]
			flag[n] = left + right + 1
			flag[n-left] = flag[n]
			flag[n+right] = flag[n]
			if flag[n] > res {
				res = flag[n]
			}
		}
	}
	return res
}
