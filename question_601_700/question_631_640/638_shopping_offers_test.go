package question_631_640

import (
	"testing"
)

func Test_shoppingOffers(t *testing.T) {
	tests := []struct {
		price   []int
		special [][]int
		needs   []int
		want    int
	}{
		{[]int{2, 5}, [][]int{{3, 0, 5}, {1, 2, 10}}, []int{3, 2}, 0},
		//{[]int{2, 3, 4}, [][]int{{1, 1, 0, 4}, {2, 2, 1, 9}}, []int{1, 2, 1}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shoppingOffers(tt.price, tt.special, tt.needs); got != tt.want {
				t.Errorf("shoppingOffers() = %v, want %v", got, tt.want)
			}
		})
	}
}
