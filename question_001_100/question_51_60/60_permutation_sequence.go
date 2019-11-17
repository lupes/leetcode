package question_51_60

// 60. 第k个排列
// https://leetcode-cn.com/problems/permutation-sequence
// Topics: 数学 回溯算法

func getPermutation(n int, k int) string {
	if n == 0 || k == 0 {
		return ""
	}
	nums := 1
	for i := 2; i < n; i++ {
		nums *= i
	}
	if k > nums*n {
		return ""
	}
	var flag = make([]bool, n)
	var res = ""
	for n > 1 {
		i, d := k/nums, k%nums
		if k%nums != 0 {
			i++
		} else {
			d = nums
		}
		res += getKNum(flag, i)
		nums, n, k = nums/(n-1), n-1, d
	}
	return res + getKNum(flag, 1)
}

func getKNum(flag []bool, k int) string {
	var j = 0
	for i, b := range flag {
		if !b {
			j++
			if j == k {
				flag[i] = true
				return string('0' + 1 + i)
			}
		}
	}
	return ""
}
