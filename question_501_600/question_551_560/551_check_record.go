package question_551_560

// 551. 学生出勤记录 I
// https://leetcode-cn.com/problems/student-attendance-record-i/

func checkRecord(s string) bool {
	preL, hasA := 0, false
	for _, c := range s {
		switch {
		case c == 'A' && hasA:
			return false
		case c == 'A':
			hasA = true
			preL = 0
		case c == 'L' && preL == 2:
			return false
		case c == 'L':
			preL++
		case c == 'P':
			preL = 0
		}
	}
	return true
}
