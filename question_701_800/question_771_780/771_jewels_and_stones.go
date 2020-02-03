package question_771_780

// 771. 宝石与石头
// https://leetcode-cn.com/problems/jewels-and-stones
// Topics: 哈希表

func numJewelsInStones(J string, S string) int {
	flag, res := make(map[int32]struct{}), 0
	for _, c := range J {
		flag[c] = struct{}{}
	}

	for _, c := range S {
		if _, ok := flag[c]; ok {
			res++
		}
	}
	return res
}
