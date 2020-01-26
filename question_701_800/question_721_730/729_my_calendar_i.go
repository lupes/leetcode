package question_721_730

// 729. 我的日程安排表 I
// https://leetcode-cn.com/problems/my-calendar-i
// Topics: 数组

type schedule struct {
	start int
	end   int
}

type MyCalendar struct {
	schedules []schedule
}

func Constructor() MyCalendar {
	return MyCalendar{
		schedules: []schedule{{start: -1, end: 0}},
	}
}

func (this *MyCalendar) Book(start int, end int) bool {
	var i = 1
	for i = 1; i < len(this.schedules); i++ {
		schedule := this.schedules[i]
		if start >= schedule.end {
			continue
		} else if schedule.start >= end && start >= this.schedules[i-1].end {
			break
		} else if (schedule.end > start && start >= schedule.start) ||
			(schedule.end > end && end >= schedule.start) ||
			(end > schedule.start && schedule.start >= start) {
			return false
		}
	}
	newSchedules := append([]schedule{{start: start, end: end}}, this.schedules[i:]...)
	this.schedules = append(this.schedules[:i], newSchedules...)
	return true
}
