package question_591_600

// 599. 两个列表的最小索引总和
// https://leetcode-cn.com/problems/minimum-index-sum-of-two-lists/
// Topics: 哈希表

func findRestaurant(list1 []string, list2 []string) []string {
	l1, l2 := len(list1), len(list2)
	if l1 > l2 {
		list1, list2 = list2, list1
	}
	var tmp = make(map[string]int)
	for i, s := range list1 {
		tmp[s] = i
	}
	var res []string
	var sum = l1 + l2
	for j, s := range list2 {
		if i, ok := tmp[s]; ok {
			if i+j < sum {
				res = []string{s}
				sum = i + j
			} else if i+j == sum {
				res = append(res, s)
			}
		}
	}
	return res
}
