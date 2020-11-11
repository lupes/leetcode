package question_1331_1340

import (
	"time"
)

// 1360. 日期之间隔几天
// https://leetcode-cn.com/problems/number-of-days-between-two-dates/
// Topics: 日期

func daysBetweenDates(date1 string, date2 string) int {
	d1, _ := time.Parse("2006-01-02", date1)
	d2, _ := time.Parse("2006-01-02", date2)

	d := int(d2.Sub(d1).Hours() / 24)
	if d > 0 {
		return d
	}
	return -d
}
