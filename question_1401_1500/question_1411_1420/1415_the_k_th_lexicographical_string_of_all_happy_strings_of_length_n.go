package question_1411_1420

// 1415. 长度为 n 的开心字符串中字典序第 k 小的字符串
// https://leetcode-cn.com/problems/the-k-th-lexicographical-string-of-all-happy-strings-of-length-n/
// Topics: 栈 设计

func getHappyString(n int, k int) string {
	var flag, sum = make([]byte, n+1), make(map[int]int, 1)
	_, res := getHappyStringHelper(n, k, flag, sum, 1)
	return res
}

func getHappyStringHelper(n int, k int, flag []byte, sum map[int]int, i int) (bool, string) {
	if i > n {
		sum[1]++
		if sum[1] == k {
			return true, string(flag[1:])
		}
		return false, ""
	}

	for _, c := range []byte{'a', 'b', 'c'} {
		if flag[i-1] == c {
			continue
		}

		flag[i] = c
		find, res := getHappyStringHelper(n, k, flag, sum, i+1)
		if find {
			return true, res
		}
	}
	return false, ""
}
