package qustion_511_520

import "math/rand"

// 519. 随机翻转矩阵
// https://leetcode-cn.com/problems/random-flip-matrix/

type Solution struct {
	Nums int
	Cols int
	Now  int
	Map  map[int]struct{}
}

func Constructor(n_rows int, n_cols int) Solution {
	return Solution{
		Nums: n_rows * n_cols,
		Cols: n_cols,
		Now:  0,
		Map:  make(map[int]struct{}),
	}
}

func (this *Solution) Flip() []int {
	if this.Now == this.Nums {
		this.Reset()
	}
	for n := rand.Intn(this.Nums); ; {
		if _, ok := this.Map[n]; !ok {
			this.Now++
			this.Map[n] = struct{}{}
			return []int{n / this.Cols, n % this.Cols}
		}
		if n++; n == this.Nums {
			n = 0
		}
	}
}

func (this *Solution) Reset() {
	this.Map = make(map[int]struct{})
	this.Now = 0
}
