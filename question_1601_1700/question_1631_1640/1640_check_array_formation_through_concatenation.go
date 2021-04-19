package question_1631_1640

// 1640. 能否连接形成数组
// https://leetcode-cn.com/problems/check-array-formation-through-concatenation/
// Topics: 排序 数组 哈希表

func canFormArray(arr []int, pieces [][]int) bool {
	var flag = make(map[int]int)
	for i, price := range pieces {
		flag[price[0]] = i
	}

	for i := 0; i < len(arr); {
		j, ok := flag[arr[i]]
		if !ok {
			return false
		}
		for _, n := range pieces[j] {
			if n != arr[i] {
				return false
			}
			i++
		}
	}
	return true
}
