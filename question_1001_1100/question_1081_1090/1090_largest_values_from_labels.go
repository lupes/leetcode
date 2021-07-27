package question_1081_1090

import (
	"sort"
)

// 1090. 受标签影响的最大值
// https://leetcode-cn.com/problems/largest-values-from-labels
// Topics: 贪心算法 哈希表

func largestValsFromLabels(values []int, labels []int, num_wanted int, use_limit int) int {
	var tmp = make([][2]int, 0, len(values))
	for i := range values {
		tmp = append(tmp, [2]int{values[i], labels[i]})
	}
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][0] > tmp[j][0]
	})
	var max, now, flag = 0, 0, make(map[int]int, 0)
	for i := 0; i < len(values) && now < num_wanted; i++ {
		if n := flag[tmp[i][1]]; n < use_limit {
			max += tmp[i][0]
			now++
			flag[tmp[i][1]]++
		}
	}
	return max
}
