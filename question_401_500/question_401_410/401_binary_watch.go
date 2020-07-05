package question_401_410

// 401. 二进制手表
// https://leetcode-cn.com/problems/binary-watch
// Topics: 位运算 回溯算法

func readBinaryWatch(num int) []string {
	var tmp = make([]int, 10)
	var res = make([]string, 0)
	readBinaryWatchHelper(&res, tmp, num, 0)
	return res
}

func readBinaryWatchHelper(res *[]string, tmp []int, num, start int) {
	if num == 0 {
		addTime(res, tmp)
	}
	num -= 1
	for i := start; i < 10; i++ {
		tmp[i] = 1
		readBinaryWatchHelper(res, tmp, num, i+1)
		tmp[i] = 0
	}
}

func addTime(res *[]string, tmp []int) {
	hour := tmp[0]*8 + tmp[1]*4 + tmp[2]*2 + tmp[3]
	minute := tmp[4]*32 + tmp[5]*16 + tmp[6]*8 + tmp[7]*4 + tmp[8]*2 + tmp[9]
	if hour < 12 && minute < 60 {
		a, b, c, d := hour/10, hour%10, minute/10, minute%10
		s := string('0'+a) + string('0'+b) + ":" + string('0'+c) + string('0'+d)
		if a == 0 {
			s = s[1:]
		}
		*res = append(*res, s)
	}
}
