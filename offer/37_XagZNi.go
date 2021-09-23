package offer

// 剑指 Offer II 037. 小行星碰撞
// https://leetcode-cn.com/problems/XagZNi/
// Topics: 栈 数组

func asteroidCollision(asteroids []int) []int {
	var stack []int
	for _, asteroid := range asteroids {
		if asteroid > 0 {
			stack = append(stack, asteroid)
		} else {
			for len(stack) > 0 && stack[len(stack)-1] > 0 && asteroid != 0 {
				if stack[len(stack)-1] > -asteroid {
					asteroid = 0
				} else if stack[len(stack)-1] < -asteroid {
					stack = stack[:len(stack)-1]
				} else if stack[len(stack)-1] == -asteroid {
					asteroid = 0
					stack = stack[:len(stack)-1]
				}
			}
			if asteroid != 0 {
				stack = append(stack, asteroid)
			}
		}
	}
	return stack
}
