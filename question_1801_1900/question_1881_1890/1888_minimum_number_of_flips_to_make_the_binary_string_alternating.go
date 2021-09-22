package question_1881_1890

// 1888. 使二进制字符串字符交替的最少反转次数
// https://leetcode-cn.com/problems/minimum-number-of-flips-to-make-the-binary-string-alternating/
// Topics: 贪心算法 字符串

func minFlips(s string) int {
	var size = len(s)
	var flag = make([][2]int, len(s)+1)
	for i, c := range s {
		flag[i+1][1] = flag[i][1]
		flag[i+1][0] = flag[i][0]
		if i%2 == 0 && c == '0' {
			flag[i+1][1] += 1
		} else if i%2 == 0 && c == '1' {
			flag[i+1][0] += 1
		} else if c == '0' {
			flag[i+1][0] += 1
		} else {
			flag[i+1][1] += 1
		}
	}
	if flag[size][0] == 0 || flag[size][1] == 0 {
		return 0
	}
	var res = flag[size][0]
	if flag[size][0] > flag[size][1] {
		res = flag[size][1]
	}
	if size%2 == 0 {
		return res
	}
	for i := 1; i < size; i++ {
		tmp1 := flag[i][0] + flag[size][1] - flag[i][1]
		tmp2 := flag[i][1] + flag[size][0] - flag[i][0]
		if tmp1 > tmp2 {
			tmp1 = tmp2
		}
		if tmp1 < res {
			res = tmp1
		}
	}
	return res
}
