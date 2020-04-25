package question_1001_1010

// 1004. 最大连续1的个数 III
// https://leetcode-cn.com/problems/max-consecutive-ones-iii
// Topics: 双指针 None

func longestOnes(A []int, K int) int {
	left, right, tmp, res := 0, 0, 0, 0
	for len(A) > left && len(A) > right {
		if tmp <= K {
			if A[right] == 0 {
				tmp++
			}
			if tmp <= K && right-left+1 > res {
				res = right - left + 1
			}
			right++
		} else {
			if A[left] == 0 {
				tmp--
			}
			left++
		}
		//fmt.Printf("left:%d right:%d tmp:%d res:%d\n", left, right, tmp, res)
	}
	return res
}
