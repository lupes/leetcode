package question_1091_1100

// 1093. 大样本统计
// https://leetcode-cn.com/problems/statistics-from-a-large-sample
// Topics: 数学 双指针

func sampleStats(count []int) []float64 {
	var maxZ = 0
	var index = 0
	var sum = 0
	var num = 0
	var min, max = -1, -1
	for i, n := range count {
		if n == 0 {
			continue
		}
		if min == -1 {
			min = i
		}
		max = i
		sum += i * n
		if n > maxZ {
			maxZ = n
			index = i
		}
		num += n
	}
	an, bn := num/2+1, num/2+1
	if num%2 == 0 {
		an, bn = num/2, num/2+1
	}

	a, b := 0, 0
	for i, n := range count {
		if n == 0 {
			continue
		}
		if an-n > 0 {
			an -= n
		} else if an > 0 {
			an -= n
			a = i
		}
		if bn-n > 0 {
			bn -= n
		} else {
			b = i
			break
		}
	}
	return []float64{float64(min), float64(max), float64(sum) / float64(num), (float64(a) + float64(b)) / 2, float64(index)}
}
