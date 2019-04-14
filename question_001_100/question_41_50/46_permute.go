package question_41_50

func permute(nums []int) [][]int {
	flag := make([]bool, len(nums))
	return dfs1(nums, flag)
}

func dfs1(nums []int, flag []bool) [][]int {
	var res [][]int
	for i, num := range nums {
		if !flag[i] {
			var newFlag = make([]bool, len(flag))
			copy(newFlag, flag)
			newFlag[i] = true
			tmp := dfs1(nums, newFlag)
			newArr := []int{num}
			if tmp == nil {
				res = append(res, newArr)
			}
			for _, t := range tmp {
				r := append(newArr, t...)
				res = append(res, r)
			}
		}
	}
	return res
}
