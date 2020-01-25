package question_711_720

// 717. 1比特与2比特字符
// https://leetcode-cn.com/problems/1-bit-and-2-bit-characters
// Topics: 数组

func isOneBitCharacter(bits []int) bool {
	i, l := 0, len(bits)
	for i = 0; i < l-1; i++ {
		if bits[i] == 1 {
			i++
		}
	}
	return i == l-1
}
