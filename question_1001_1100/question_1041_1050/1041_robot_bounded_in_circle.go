package question_1041_1050

// 1041. 困于环中的机器人
// https://leetcode-cn.com/problems/robot-bounded-in-circle
// Topics: 数学

func isRobotBounded(instructions string) bool {
	x, y, t := 0, 0, 4
	for i := 0; i < 4; i++ {
		for _, c := range instructions {
			if c == 'G' {
				if t == 1 {
					x += 1
				} else if t == 2 {
					y -= 1
				} else if t == 3 {
					x -= 1
				} else if t == 4 {
					y += 1
				}
			} else if c == 'L' {
				t -= 1
				if t == 0 {
					t = 4
				}
			} else if c == 'R' {
				t += 1
				if t == 5 {
					t = 1
				}
			}
		}
	}
	if x == 0 && y == 0 && t == 4 {
		return true
	}
	return false
}
