package question_571_580

// 575. 分糖果
// https://leetcode-cn.com/problems/distribute-candies
// Topics: 哈希表

func distributeCandies(candies []int) int {
	var tmp = make(map[int]struct{})
	var r = len(candies) / 2
	var res = 0
	for _, c := range candies {
		if _, ok := tmp[c]; !ok {
			res++
			tmp[c] = struct{}{}
			if res == r {
				return res
			}
		}
	}
	return res
}
