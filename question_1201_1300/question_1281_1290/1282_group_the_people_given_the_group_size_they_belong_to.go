package question_1281_1290

// 1282. 用户分组
// https://leetcode-cn.com/problems/group-the-people-given-the-group-size-they-belong-to/
// Topics: 贪心算法

func groupThePeople(groupSizes []int) [][]int {
	var res [][]int
	var flag = make(map[int][]int, 0)
	for i, size := range groupSizes {
		flag[size] = append(flag[size], i)
		if len(flag[size]) == size {
			res = append(res, flag[size])
			flag[size] = nil
		}
	}
	return res
}
