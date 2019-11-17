package question_341_350

// 349. 两个数组的交集
// https://leetcode-cn.com/problems/intersection-of-two-arrays/
// Topics: 排序 哈希表 双指针 二分查找

func intersection(nums1 []int, nums2 []int) []int {
	var m = make(map[int]struct{}, 0)
	var r = make(map[int]struct{}, 0)
	for _, n := range nums1 {
		m[n] = struct{}{}
	}
	for _, n := range nums2 {
		_, ok := m[n]
		if ok {
			r[n] = struct{}{}
		}
	}
	var res []int
	for k, _ := range r {
		res = append(res, k)
	}
	return res
}
