package question_1681_1690

// 1689. 十-二进制数的最少数目
// https://leetcode-cn.com/problems/partitioning-into-minimum-number-of-deci-binary-numbers/
// Topics: 贪心算法

func minPartitions(n string) int {
	var res = uint8(0)
	for i := 0; i < len(n) && res != 9; i++ {
		if n[i]-'0' > res {
			res = n[i] - '0'
		}
	}
	return int(res)
}
