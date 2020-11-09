package question_1331_1340

import (
	"testing"
)

func TestCashier(t *testing.T) {
	this := Constructor2(3, 50, []int{1, 2, 3, 4, 5, 6, 7}, []int{100, 200, 300, 400, 300, 200, 100})
	if got, want := this.GetBill([]int{1, 2}, []int{1, 2}), 500.0; got != want {
		t.Errorf("GetBill() = %v, want %v", got, want)
	}

	if got, want := this.GetBill([]int{3, 7}, []int{10, 10}), 4000.0; got != want {
		t.Errorf("GetBill() = %v, want %v", got, want)
	}

	if got, want := this.GetBill([]int{1, 2, 3, 4, 5, 6, 7}, []int{1, 1, 1, 1, 1, 1, 1}), 800.0; got != want {
		t.Errorf("GetBill() = %v, want %v", got, want)
	}
}
