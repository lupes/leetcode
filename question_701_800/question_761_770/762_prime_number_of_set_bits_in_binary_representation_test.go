package question_761_770

import (
	"testing"
)

func Test_countPrimeSetBits(t *testing.T) {
	tests := []struct {
		L    int
		R    int
		want int
	}{
		{6, 10, 4},
		{10, 15, 5},
		{1, 1000000, 322931},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countPrimeSetBits(tt.L, tt.R); got != tt.want {
				t.Errorf("countPrimeSetBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
