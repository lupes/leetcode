package common

// 归并排序
func mergeSort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}
	left := mergeSort(nums[:len(nums)/2])
	right := mergeSort(nums[len(nums)/2:])

	var res = make([]int, 0, len(nums))
	for l, r := 0, 0; l < len(left) || r < len(right); {
		if l == len(left) {
			res = append(res, right[r])
			r++
		} else if r == len(right) {
			res = append(res, left[l])
			l++
		} else if left[l] > right[r] {
			res = append(res, right[r])
			r++
		} else {
			res = append(res, left[l])
			l++
		}
	}
	return res
}

func resortList(root *ListNode) *ListNode {
	if root == nil {
		return nil
	}
	head, next := &ListNode{Next: root}, root.Next
	root.Next = nil
	for next != nil {
		head.Next, next.Next, next = next, head.Next, next.Next
	}
	return head.Next
}
