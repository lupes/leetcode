package question_121_130

// 128. 最长连续序列
// https://leetcode-cn.com/problems/longest-consecutive-sequence
// Topics: 并查集 数组

func longestConsecutive(nums []int) int {
	var res, min, max, flag = 0, make(map[int]int), make(map[int]int), make(map[int]struct{})
	for _, n := range nums {
		if _, ok := flag[n]; !ok {
			start, end := n, n+1
			if t, ok := min[n+1]; ok {
				delete(min, n+1)
				delete(max, t)
				end = t
			}
			if t, ok := max[n]; ok {
				delete(max, n)
				delete(min, t)
				start = t
			}
			min[start] = end
			max[end] = start
			flag[n] = struct{}{}
			if end-start > res {
				res = end - start
			}
		}
	}
	return res
}
