package question_831_840

// 838. 推多米诺
// https://leetcode-cn.com/problems/push-dominoes
// Topics: 双指针 动态规划

func pushDominoes(dominoes string) string {
	var res = []byte("L" + dominoes + "R")
	left, right := -1, -1
	for i := 0; i < len(res); i++ {
		if res[i] == '.' {
			if left == -1 {
				left = i - 1
			}
			right = i + 1
		} else {
			if right != -1 {
				if res[left] == 'R' && res[right] == 'L' {
					for j := 0; j < (right-left-1)/2; j++ {
						res[left+j+1] = 'R'
						res[right-j-1] = 'L'
					}
				} else if res[left] == res[right] {
					for j := left + 1; j < right; j++ {
						res[j] = res[left]
					}
				}
			}
			left, right = -1, -1
		}
	}
	return string(res[1 : len(res)-1])
}
