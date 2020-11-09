package question_1331_1340

// 1357. 每隔 n 个顾客打折
// https://leetcode-cn.com/problems/apply-discount-every-n-orders/
// Topics: 设计

type Cashier struct {
	N        int
	Now      int
	Discount int
	Flag     map[int]int
}

func Constructor2(n int, discount int, products []int, prices []int) Cashier {
	var cashier = Cashier{
		N:        n,
		Now:      0,
		Discount: discount,
		Flag:     make(map[int]int),
	}
	for i, product := range products {
		cashier.Flag[product] = prices[i]
	}
	return cashier
}

func (this *Cashier) GetBill(product []int, amount []int) float64 {
	var total, res = 0, 0.0
	for i, p := range product {
		total += this.Flag[p] * amount[i]
	}
	this.Now++
	this.Now %= this.N
	res = float64(total)
	if this.Now == 0 {
		res = res * float64(100-this.Discount) / 100
	}
	return res
}
