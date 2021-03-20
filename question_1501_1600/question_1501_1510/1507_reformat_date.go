package question_1501_1510

// 1507. 转变日期格式
// https://leetcode-cn.com/problems/reformat-date/
// Topics: 字符串

var months = map[string]string{
	"Jan": "01", "Feb": "02", "Mar": "03", "Apr": "04", "May": "05", "Jun": "06",
	"Jul": "07", "Aug": "08", "Sep": "09", "Oct": "10", "Nov": "11", "Dec": "12",
}

func reformatDate(date string) string {
	if len(date) == 12 {
		date = "0" + date
	}
	return date[9:13] + "-" + months[date[5:8]] + "-" + date[0:2]
}
