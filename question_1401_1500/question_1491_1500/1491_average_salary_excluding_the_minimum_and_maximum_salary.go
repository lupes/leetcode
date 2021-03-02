package question_1491_1500

// 1491. 去掉最低工资和最高工资后的工资平均值
// https://leetcode-cn.com/problems/average-salary-excluding-the-minimum-and-maximum-salary/
// Topics: 排序 数组

func average(salary []int) float64 {
	min, max := salary[0], salary[1]
	if salary[0] > salary[1] {
		min, max = salary[1], salary[0]
	}
	var sum = 0
	for i := 2; i < len(salary); i++ {
		if salary[i] > max {
			sum += max
			max = salary[i]
		} else if salary[i] < min {
			sum += min
			min = salary[i]
		} else {
			sum += salary[i]
		}
	}
	return float64(sum) / float64(len(salary)-2)
}
