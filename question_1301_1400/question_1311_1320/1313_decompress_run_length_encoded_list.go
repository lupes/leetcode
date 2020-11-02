package question_1311_1320

// 1313. 解压缩编码列表
// https://leetcode-cn.com/problems/decompress-run-length-encoded-list/
// Topics: 数组

func decompressRLElist(nums []int) []int {
	var res []int
	for i := 0; i < len(nums)/2; i++ {
		for j := 0; j < nums[i*2]; j++ {
			res = append(res, nums[i*2+1])
		}
	}
	return res
}
