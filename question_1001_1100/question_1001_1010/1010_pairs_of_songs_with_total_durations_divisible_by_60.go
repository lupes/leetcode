package question_1001_1010

// 1010. 总持续时间可被 60 整除的歌曲
// https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60
// Topics: 数组

func numPairsDivisibleBy60(time []int) int {
	var flag, res = make([]int, 60), 0
	for _, n := range time {
		t := n % 60
		if t == 0 {
			res += flag[0]
		} else {
			res += flag[60-t]
		}
		flag[t]++
	}
	return res
}
