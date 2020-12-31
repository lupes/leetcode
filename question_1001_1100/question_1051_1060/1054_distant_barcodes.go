package question_1051_1060

import (
	"sort"
)

// 1054. 距离相等的条形码
// https://leetcode-cn.com/problems/distant-barcodes
// Topics: 堆 排序

type Barcode struct {
	barcode int
	num     int
}

func rearrangeBarcodes(barcodes []int) []int {
	var flag = make(map[int]int)
	for _, barcode := range barcodes {
		flag[barcode]++
	}

	var codes = make([]Barcode, 0, len(flag))
	for k, v := range flag {
		codes = append(codes, Barcode{
			barcode: k,
			num:     v,
		})
	}

	sort.Slice(codes, func(i, j int) bool {
		return codes[i].num > codes[j].num
	})

	var res []int
	for i := 0; i < codes[0].num; i++ {
		res = append(res, codes[0].barcode)
	}
	codes = codes[1:]

	for {
		var next []int
		for i, n := range res {
			next = append(next, n)
			if len(codes) > 0 {
				next = append(next, codes[0].barcode)
				if codes[0].num == 1 {
					codes = codes[1:]
				} else {
					codes[0].num--
				}
			} else {
				next = append(next, res[i+1:]...)
				return next
			}
		}
		res = next
	}
}
