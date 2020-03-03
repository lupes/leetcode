package question_861_870

// 869. 重新排序得到 2 的幂
// https://leetcode-cn.com/problems/reordered-power-of-2
// Topics: 数学

var flag = map[int]struct{}{
	1: {}, 2: {}, 4: {}, 8: {}, 61: {}, 32: {}, 64: {}, 821: {}, 652: {}, 521: {},
	4210: {}, 8420: {}, 9640: {}, 9821: {}, 86431: {}, 87632: {}, 66553: {}, 732110: {},
	644221: {}, 885422: {}, 8765410: {}, 9752210: {}, 9444310: {}, 8888630: {}, 77766211: {},
	55443332: {}, 88766410: {}, 877432211: {}, 866554432: {}, 987653210: {},
}

func reorderedPowerOf2(N int) bool {
	var f = [10]byte{}
	for t := N; t > 0; t /= 10 {
		f[t%10]++
	}

	var tmp = 0
	for i := 9; i >= 0; i-- {
		for j := byte(0); j < f[i]; j++ {
			tmp = tmp*10 + i
		}
	}
	_, ok := flag[tmp]
	return ok
}
