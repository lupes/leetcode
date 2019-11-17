package question_41_50

// 45. 跳跃游戏 II
// https://leetcode-cn.com/problems/jump-game-ii
// Topics: 贪心算法 数组

func jump(nums []int) int {
	size := len(nums)
	if size == 1 {
		return 0
	}
	var num, res, max, tmp, next, i, j int
	for i = 0; i < size; {
		res += 1
		num = nums[i]
		if num+i+1 >= size {
			return res
		}
		max = 0
		for j = 1; j <= num; j++ {
			tmp = j + nums[j+i]
			if tmp > max {
				max = tmp
				next = i + j
			}
		}
		i = next
	}
	return 0
}
