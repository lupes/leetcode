package question_491_500

// 496. 下一个更大元素 I
// https://leetcode-cn.com/problems/next-greater-element-i/
// Topics: 栈

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var res = make([]int, len(nums1))
	var l = len(nums2)
	for i, n1 := range nums1 {
		res[i] = -1
		for j := l - 1; j > 0 && nums2[j] != n1; j-- {
			if nums2[j] > n1 {
				res[i] = nums2[j]
			}
		}
	}
	return res
}
