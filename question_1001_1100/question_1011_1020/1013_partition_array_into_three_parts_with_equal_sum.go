package question_1011_1020

// 1013. 将数组分成和相等的三个部分
// https://leetcode-cn.com/problems/partition-array-into-three-parts-with-equal-sum
// Topics: 数组

func canThreePartsEqualSum(A []int) bool {
	sum, tmp, num := 0, 0, 0
	for _, n := range A {
		sum += n
	}
	if sum%3 != 0 {
		return false
	}
	sum = sum / 3
	for _, n := range A {
		tmp += n
		if tmp == sum {
			num, tmp = num+1, 0
		}
	}
	if sum == 0 && num >= 3 {
		return true
	} else if num == 3 {
		return true
	}
	return false
}
