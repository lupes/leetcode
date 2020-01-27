package question_731_740

// 735. 行星碰撞
// https://leetcode-cn.com/problems/asteroid-collision
// Topics: 栈

func asteroidCollision(asteroids []int) []int {
	var res []int
	for i := 0; i < len(asteroids); i++ {
		now := asteroids[i]
		if len(res) == 0 || now > 0 || (now < 0 && res[len(res)-1] < 0) {
			res = append(res, now)
		} else if now < 0 && now*-1 == res[len(res)-1] {
			res = res[:len(res)-1]
		} else if now < 0 && float32(now)/float32(res[len(res)-1]) < -1 {
			res = res[:len(res)-1]
			i--
		}
	}
	return res
}
