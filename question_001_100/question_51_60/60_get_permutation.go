package question_51_60

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
	return getPermutationDfs(flag, nums, n, k)
}

func getPermutationDfs(flag []bool, nums, n, k int) string {
	if n == 1 {
		return getKNum(flag, 1)
	}
	i := k / nums
	d := k % nums
	if k%nums != 0 {
		i++
	} else {
		d = nums
	}
	return getKNum(flag, i) + getPermutationDfs(flag, nums/(n-1), n-1, d)
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
