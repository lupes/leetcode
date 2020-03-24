package question_891_900

// 900. RLE 迭代器
// https://leetcode-cn.com/problems/rle-iterator
// Topics: 数组

type RLEIterator struct {
	A []int
}

func Constructor(A []int) RLEIterator {
	return RLEIterator{
		A,
	}
}

func (this *RLEIterator) Next(n int) int {
	for len(this.A) > 0 {
		if this.A[0] == 0 {
			this.A = this.A[2:]
		} else if this.A[0] > n {
			this.A[0] -= n
			return this.A[1]
		} else if this.A[0] == n {
			tmp := this.A[1]
			this.A = this.A[2:]
			return tmp
		} else {
			n -= this.A[0]
			this.A = this.A[2:]
		}
	}
	return -1
}

/**
 * Your RLEIterator object will be instantiated and called as such:
 * obj := Constructor(A);
 * param_1 := obj.Next(n);
 */
