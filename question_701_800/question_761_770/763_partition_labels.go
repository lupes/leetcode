package question_761_770

// 763. 划分字母区间
// https://leetcode-cn.com/problems/partition-labels
// Topics: 贪心算法 双指针

func partitionLabels(S string) []int {
	var sumMap = ['z' + 1]int{}
	for i := range S {
		sumMap[S[i]] = i
	}

	var res []int
	var start, end = 0, -1
	for end+1 < len(S) {
		start, end = end+1, sumMap[S[end+1]]
		for i := start; i < end; i++ {
			if sumMap[S[i]] > end {
				end = sumMap[S[i]]
			}
		}
		res = append(res, end-start+1)
	}

	return res
}
