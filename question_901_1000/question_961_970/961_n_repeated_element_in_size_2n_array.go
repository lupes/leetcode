package question_961_970

// 961. 重复 N 次的元素
// https://leetcode-cn.com/problems/n-repeated-element-in-size-2n-array
// Topics: 哈希表

func repeatedNTimes(A []int) int {
	var flag = make(map[int]int, 0)
	for _, n := range A {
		flag[n]++
		if flag[n] == 2 {
			return n
		}
	}
	return 0
}
