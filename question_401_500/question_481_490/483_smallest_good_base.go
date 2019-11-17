package question_481_490

import (
	"fmt"
	"strconv"
)

// Topics: 数学 二分查找

func smallestGoodBase(n string) string {
	num := 0
	for _, c := range n {
		num = num*10 + int(c-'0')
	}
	fmt.Printf("%#v\n", num)
	for k := 2; k < num-1; k++ {
		fmt.Printf("k %#v\n", k)
		for tmp := num; tmp%k == 1; tmp /= k {
			fmt.Printf("tmp %#v\n", tmp)
			if tmp == 1 {
				return strconv.Itoa(k)
			}
		}
	}

	return ""
}
