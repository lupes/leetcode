package question_1151_1160

// 1154. 一年中的第几天
// https://leetcode-cn.com/problems/day-of-the-year
// Topics: 数学

var days = [2][12]int{
	{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
	{31, 29, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31},
}

func dayOfYear(date string) int {
	year := int(date[0]-'0')*1000 + int(date[1]-'0')*100 + int(date[2]-'0')*10 + int(date[3]-'0')
	month := int(date[5]-'0')*10 + int(date[6]-'0')
	day := int(date[8]-'0')*10 + int(date[9]-'0')
	res, flag := day, 0
	if year%4 == 0 && year%100 != 0 {
		flag = 1
	}
	for i := 0; i < month-1; i++ {
		res += days[flag][i]
	}
	return res
}
