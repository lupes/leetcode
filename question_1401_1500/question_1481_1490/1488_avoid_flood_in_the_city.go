package question_1481_1490

// 1488. 避免洪水泛滥
// https://leetcode-cn.com/problems/avoid-flood-in-the-city/
// Topics: 哈希表 数组

func avoidFlood(rains []int) []int {
	var flag = make(map[int]int)
	var res = make([]int, len(rains))
	for i, n := range rains {
		if n == 0 {
			res[i] = -2
		} else if n > 0 {
			res[i] = -1
			if t, ok := flag[n]; ok {
				success := false
				for j := t + 1; j < i; j++ {
					if res[j] == -2 {
						res[j] = n
						success = true
						break
					}
				}
				if !success {
					return nil
				}
			}
			flag[n] = i
		}
	}

	for i, n := range res {
		if n == -2 {
			res[i] = 1
		}
	}

	return res
}
