package question_601_610

// 611. 有效三角形的个数
// https://leetcode-cn.com/problems/valid-triangle-number/

func triangleNumber(nums []int) int {
	if len(nums) < 3 {
		return 0
	}
	var res int
	for i, a := range nums[:len(nums)-2] {
		for j, b := range nums[i+1 : len(nums)-1] {
			for _, c := range nums[i+j+2:] {
				if a+b > c && a+c > b && b+c > a {
					res++
				}
			}
		}
	}
	return res
}
