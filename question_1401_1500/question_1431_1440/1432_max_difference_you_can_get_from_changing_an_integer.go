package question_1431_1440

// 1432. 改变一个整数能得到的最大差值
// https://leetcode-cn.com/problems/max-difference-you-can-get-from-changing-an-integer/
// Topics: 字符串

func maxDiff(num int) int {
	var flag []int
	for tmp := num; tmp > 0; tmp /= 10 {
		flag = append([]int{tmp % 10}, flag...)
	}

	var x, y = -1, 9
	var maxFlag = make([]int, len(flag))
	for i := 0; i < len(flag); i++ {
		if (flag[i] != 9 && x == -1) || (x == flag[i]) {
			x, maxFlag[i] = flag[i], y
		} else {
			maxFlag[i] = flag[i]
		}
	}

	x, y = -1, 0
	var minFlag = make([]int, len(flag))
	for i := 0; i < len(flag); i++ {
		if i == 0 && flag[i] > 1 {
			x, y, minFlag[i] = flag[i], 1, 1
		} else if (x == -1 && flag[i] > 1) || (flag[i] == x) {
			x, minFlag[i] = flag[i], y
		} else {
			minFlag[i] = flag[i]
		}
	}

	var max, min = 0, 0
	for i := 0; i < len(minFlag); i++ {
		max = max*10 + maxFlag[i]
		min = min*10 + minFlag[i]
	}

	return max - min
}
