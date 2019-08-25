package question_231_240

// 234. 回文链表
// https://leetcode-cn.com/problems/palindrome-linked-list/

func isPalindrome(head *ListNode) bool {
	var arr []*ListNode
	for t := head; t != nil; t = t.Next {
		arr = append(arr, t)
	}

	for l, r := 0, len(arr)-1; r > l; l, r = l+1, r-1 {
		if arr[l].Val != arr[r].Val {
			return false
		}
	}
	return true
}
