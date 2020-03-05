package question_861_870

// 868. 二进制间距
// https://leetcode-cn.com/problems/binary-gap
// Topics: 数学

func binaryGap(N int) int {
	max, last, now := 0, -1, 0
	for n := N; n > 0; n /= 2 {
		now++
		if n%2 == 1 {
			if last != -1 && now-last > max {
				max = now - last
			}
			last = now
		}
	}
	return max
}
