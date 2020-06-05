package question_1181_1190

// 1185. 一周中的第几天
// https://leetcode-cn.com/problems/day-of-the-week
// Topics: 数组

var weekdays = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

var days = [2][12]int{
	{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
}

func dayOfTheWeek(day int, month int, year int) string {
	num, flag := 0, 0
	for i := 1971; i < year; i++ {
		num += 365
		if (i%4 == 0 && i%100 != 0) || i%400 == 0 {
			num++
		}
	}

	if year%4 == 0 && year%100 != 0 {
		flag = 1
	}
	for i := 0; i < month-1; i++ {
		num += days[flag][i]
	}

	return weekdays[(num+day+3)%7]
}
