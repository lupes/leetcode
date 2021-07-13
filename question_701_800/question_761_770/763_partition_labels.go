package question_761_770

// 763. 划分字母区间
// https://leetcode-cn.com/problems/partition-labels
// Topics: 贪心算法 双指针

func partitionLabels(S string) []int {
	var flag = ['z' + 1]int{}
	for i := range S {
		flag[S[i]] = i
	}

	var res []int
	for start, end := 0, 0; start < len(S); start = end + 1 {
		end = flag[S[start]]
		for i := start; i < end; i++ {
			if flag[S[i]] > end {
				end = flag[S[i]]
			}
		}
		res = append(res, end-start+1)
	}

	return res
}
