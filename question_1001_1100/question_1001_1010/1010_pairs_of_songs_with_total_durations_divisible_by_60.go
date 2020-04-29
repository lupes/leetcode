package question_1001_1010

// 1010. 总持续时间可被 60 整除的歌曲
// https://leetcode-cn.com/problems/pairs-of-songs-with-total-durations-divisible-by-60
// Topics: 数组

func numPairsDivisibleBy60(time []int) int {
	var flag = make([][]int, 60)
	for i, n := range time {
		flag[n%60] = append(flag[n%60], i)
	}
	return 0
}
