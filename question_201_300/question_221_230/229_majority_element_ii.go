package question_221_230

// 229. 求众数 II
// https://leetcode-cn.com/problems/majority-element-ii/
// Topics: 数组

func majorityElement(nums []int) []int {
	var a, b, an, bn = 0, 0, 0, 0
	for _, n := range nums {
		if an > 0 && n == a {
			an++
		} else if bn > 0 && n == b {
			bn++
		} else if an == 0 {
			a, an = n, 1
		} else if bn == 0 {
			b, bn = n, 1
		} else {
			an, bn = an-1, bn-1
		}
	}
	an, bn = 0, 0
	for _, n := range nums {
		if n == a {
			an++
		} else if n == b {
			bn++
		}
	}
	var res []int
	if an > len(nums)/3 {
		res = append(res, a)
	}
	if bn > len(nums)/3 {
		res = append(res, b)
	}
	return res
}
