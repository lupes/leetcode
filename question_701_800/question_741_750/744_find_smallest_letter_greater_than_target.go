package question_741_750

// 744. 寻找比目标字母大的最小字母
// https://leetcode-cn.com/problems/find-smallest-letter-greater-than-target
// Topics: 二分查找

func nextGreatestLetter(letters []byte, target byte) byte {
	left, right, c := 0, len(letters), 0
	if letters[right-1] <= target {
		return letters[0]
	}
	for right > left {
		c = (left + right) / 2
		if letters[c] > target {
			right = c
		} else {
			left = c + 1
		}
	}
	return letters[left]
}
