package question_131_140

// 135. 分发糖果
// https://leetcode-cn.com/problems/candy
// Topics: 贪心算法

func candy(ratings []int) int {
	if len(ratings) == 0 {
		return 0
	}
	var flag, l, res = make([]int, len(ratings)), len(ratings), 0
	flag[0] = 1
	for i := 1; i < l; i++ {
		if ratings[i] > ratings[i-1] {
			flag[i] = flag[i-1] + 1
		} else {
			flag[i] = 1
		}
	}
	res += flag[l-1]
	for i := l - 2; i >= 0; i-- {
		if ratings[i] > ratings[i+1] && flag[i] <= flag[i+1] {
			flag[i] = flag[i+1] + 1
		}
		res += flag[i]
	}
	return res
}
