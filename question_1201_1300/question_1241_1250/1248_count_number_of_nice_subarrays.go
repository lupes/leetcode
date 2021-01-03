package question_1241_1250

// 1248. 统计「优美子数组」
// https://leetcode-cn.com/problems/count-number-of-nice-subarrays
// Topics: 双指针

func numberOfSubarrays(nums []int, k int) int {
	var left, right, lnum, rnum, num, res int
	for right < len(nums) || num == k {
		if num < k && right < len(nums) {
			if nums[right]%2 == 1 {
				num++
				if num == k {
					res += lnum * rnum
					lnum, rnum = 1, 1
				}
			} else {
				rnum++
			}
			right++
		} else if num == k {
			if nums[left]%2 == 1 {
				num--
			} else {
				lnum++
			}
			left++
		}
	}

	return res + lnum*rnum
}

func numberOfSubarrays2(nums []int, k int) int {
	var tmp, res = []int{0}, 0
	for i, n := range nums {
		if n%2 == 1 {
			tmp = append(tmp, i+1)
		}
	}
	tmp = append(tmp, len(nums)+1)
	for i := 1; i < len(tmp)-k; i++ {
		res += (tmp[i] - tmp[i-1]) * (tmp[i+k] - tmp[i+k-1])
	}
	return res
}
