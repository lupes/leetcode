package question_311_320

import (
	"testing"
)

func Test_nthSuperUglyNumber(t *testing.T) {
	tests := []struct {
		n      int
		primes []int
		want   int
	}{
		{12, []int{2, 7, 13, 19}, 32},
		{1, []int{2, 3, 5}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := nthSuperUglyNumber(tt.n, tt.primes); got != tt.want {
				t.Errorf("nthSuperUglyNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
