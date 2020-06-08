package question_1121_1130

// 1128. 等价多米诺骨牌对的数量
// https://leetcode-cn.com/problems/number-of-equivalent-domino-pairs
// Topics: 数组

func numEquivDominoPairs(dominoes [][]int) int {
	var res, t, flag = 0, 0, make([]int, 100)
	for _, v := range dominoes {
		t = v[0]*10 + v[1]
		if v[0] > v[1] {
			t = v[1]*10 + v[0]
		}
		res += flag[t]
		flag[t]++
	}
	return res
}
