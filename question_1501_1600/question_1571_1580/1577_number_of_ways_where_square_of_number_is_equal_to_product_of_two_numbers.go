package question_1571_1580

// 1577. 数的平方等于两数乘积的方法数
// https://leetcode-cn.com/problems/number-of-ways-where-square-of-number-is-equal-to-product-of-two-numbers/
// Topics: 哈希表 数学

func numTriplets(nums1 []int, nums2 []int) int {
	flag1, flag2 := make(map[int]int), make(map[int]int)
	for _, n := range nums1 {
		flag1[n*n]++
	}

	for _, n := range nums2 {
		flag2[n*n]++
	}

	var res = 0
	for i, a := range nums1 {
		for _, c := range nums1[i+1:] {
			if n, ok := flag2[a*c]; ok {
				res += n
			}
		}
	}
	for i, a := range nums2 {
		for _, c := range nums2[i+1:] {
			if n, ok := flag1[a*c]; ok {
				res += n
			}
		}
	}
	return res
}
