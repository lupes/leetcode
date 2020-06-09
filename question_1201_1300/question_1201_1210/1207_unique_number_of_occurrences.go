package question_1201_1210

// 1207. 独一无二的出现次数
// https://leetcode-cn.com/problems/unique-number-of-occurrences
// Topics: 哈希表

func uniqueOccurrences(arr []int) bool {
	var flag1 = make(map[int]int)
	var flag2 = make(map[int]struct{})
	for _, n := range arr {
		flag1[n]++
	}
	for _, v := range flag1 {
		if _, ok := flag2[v]; ok {
			return false
		} else {
			flag2[v] = struct{}{}
		}
	}
	return true
}
