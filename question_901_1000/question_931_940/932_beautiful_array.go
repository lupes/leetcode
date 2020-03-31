package question_931_940

// 932. 漂亮数组
// https://leetcode-cn.com/problems/beautiful-array
// Topics: 分治算法

func beautifulArray(N int) []int {
	return beautifulArrayHelper(make(map[int][]int, 0), N)
}

func beautifulArrayHelper(flag map[int][]int, N int) []int {
	if res, ok := flag[N]; ok {
		return res
	}
	ans := make([]int, N)
	if N == 1 {
		ans[0] = 1
	} else {
		t := 0
		for _, x := range beautifulArrayHelper(flag, (N+1)/2) {
			ans[t] = 2*x - 1
			t++
		}
		for _, x := range beautifulArrayHelper(flag, N/2) {
			ans[t] = 2 * x
			t++
		}
	}
	flag[N] = ans
	return ans
}
