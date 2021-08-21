package question_441_450

// 443. 压缩字符串
// https://leetcode-cn.com/problems/string-compression/
// Topics: 字符串

func compress(chars []byte) int {
	tmp := append(chars, 0)
	low, fast, num := 0, 1, 1
	for ; fast < len(tmp); fast++ {
		if tmp[fast] == chars[low] {
			num++
		} else {
			if num > 1 {
				var tmp []byte
				for ; num > 0; num /= 10 {
					tmp = append([]byte{byte(num%10) + '0'}, tmp...)
				}
				for _, b := range tmp {
					low++
					chars[low] = b
				}
			}
			low++
			if low < len(chars) {
				chars[low], num = tmp[fast], 1
			}
		}
	}
	return low
}
