package question_641_650

// 645. 错误的集合
// https://leetcode-cn.com/problems/set-mismatch
// Topics: 哈希表 数学

func findErrorNums(nums []int) []int {
	var tmp = make(map[int]struct{})
	a, b := 0, 0
	for _, n := range nums {
		if _, ok := tmp[n]; ok {
			a = n
		} else {
			tmp[n] = struct{}{}
		}
	}
	for i := range nums {
		if _, ok := tmp[i+1]; !ok {
			b = i + 1
		}
	}
	return []int{a, b}
}
