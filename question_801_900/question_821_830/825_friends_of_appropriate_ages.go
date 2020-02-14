package question_821_830

// 825. 适龄的朋友
// https://leetcode-cn.com/problems/friends-of-appropriate-ages
// Topics: 数组

func numFriendRequests(ages []int) int {
	var res = 0
	var flag = make([]int, 121)
	for _, age := range ages {
		flag[age]++
	}

	for i := 120; i > 1; i-- {
		if flag[i] == 0 {
			continue
		} else if i+14 < 2*i {
			res += flag[i] * (flag[i] - 1)
		}
		for j := i - 1; j > 0; j-- {
			if i+14 >= 2*j {
				break
			}
			res += flag[j] * flag[i]
		}
	}
	return res
}
