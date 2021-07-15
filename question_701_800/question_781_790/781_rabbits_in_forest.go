package question_781_790

// 781. 森林中的兔子
// https://leetcode-cn.com/problems/rabbits-in-forest
// Topics: 哈希表 数学

func numRabbits(answers []int) int {
	var res, flag = 0, make(map[int]int)
	for _, n := range answers {
		flag[n]++
		if flag[n] == 1 {
			res += n + 1
		} else if flag[n] == n+1 {
			flag[n] = 0
		}
	}
	return res
}
