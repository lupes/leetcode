package question_1161_1170

import (
	"strconv"
	"strings"
)

// 1169. 查询无效交易
// https://leetcode-cn.com/problems/invalid-transactions
// Topics: 数组 字符串

type Transaction struct {
	index  int
	name   string
	time   int
	amount int
	city   string
}

func invalidTransactions(transactions []string) []string {
	var flag = make(map[string][]Transaction)
	var f = make(map[int]struct{})
	var res []string
	for i, transaction := range transactions {
		tmp := strings.Split(transaction, ",")
		t, _ := strconv.Atoi(tmp[1])
		a, _ := strconv.Atoi(tmp[2])
		flag[tmp[0]] = append(flag[tmp[0]], Transaction{
			index:  i,
			name:   tmp[0],
			time:   t,
			amount: a,
			city:   tmp[3],
		})
		if a > 1000 {
			f[i] = struct{}{}
		}
	}

	for _, values := range flag {
		for i := 0; i < len(values); i++ {
			for j := i + 1; j < len(values); j++ {
				if ((values[i].time >= values[j].time && values[i].time <= values[j].time+60) ||
					(values[i].time < values[j].time && values[i].time >= values[j].time-60)) &&
					values[i].city != values[j].city {
					f[values[i].index] = struct{}{}
					f[values[j].index] = struct{}{}
				}
			}
		}
	}
	for k, _ := range f {
		res = append(res, transactions[k])
	}
	return res
}
