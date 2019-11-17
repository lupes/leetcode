package question_341_350

// 350. 两个数组的交集 II
// https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/
// Topics: 排序 哈希表 双指针 二分查找

func intersect(nums1 []int, nums2 []int) []int {
	var res []int
	m := make(map[int]int)
	for _, n := range nums1 {
		m[n]++
	}
	for _, n := range nums2 {
		if t, ok := m[n]; ok && t > 0 {
			res = append(res, n)
			m[n]--
		}
	}
	return res
}
