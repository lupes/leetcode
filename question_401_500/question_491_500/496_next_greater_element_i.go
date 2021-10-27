package question_491_500

// 496. 下一个更大元素 I
// https://leetcode-cn.com/problems/next-greater-element-i/
// Topics: 栈 数组 哈希表 单调栈

func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var stack []int
	var flag = make(map[int]int, len(nums2))
	for i := len(nums2) - 1; i >= 0; i-- {
		for len(stack) > 0 && stack[len(stack)-1] < nums2[i] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			flag[nums2[i]] = -1
		} else {
			flag[nums2[i]] = stack[len(stack)-1]
		}
		stack = append(stack, nums2[i])
	}

	for i, n := range nums1 {
		nums1[i] = flag[n]
	}

	return nums1
}
