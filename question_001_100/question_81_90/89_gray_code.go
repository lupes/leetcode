package question_81_90

import "fmt"

func grayCode(n int) []int {
	var res []int
	for i := 0; i < 1<<uint(n); i++ {
		res = append(res, i^(i>>1))
	}
	return res
}

func grayCode2(n int) []int {
	flag := make([]bool, n)
	return grayCodeDfs(flag, 0)
}

func grayCodeDfs(flag []bool, s int) []int {
	var res []int
	if s == len(flag) {
		return []int{get(flag)}
	}
	res = append(res, grayCodeDfs(flag, s+1)...)

	flag[s] = true
	res = append(res, grayCodeDfs(flag, s+1)...)
	flag[s] = false
	return res
}

func get(flag []bool) int {
	var (
		s = len(flag)
		n = 0
		t = ""
	)
	for i := 0; i < s; i++ {
		if flag[i] {
			t += "1"
			n += 2 << uint(s-i-1) / 2
		} else {
			t += "0"
		}
	}
	fmt.Printf("get:%s\n", t)
	return n
}
