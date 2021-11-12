package question_371_380

import (
	"testing"
)

func Test_getMoneyAmount(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 0},
		{2, 1},
		{10, 16},
		{200, 952},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getMoneyAmount(tt.n); got != tt.want {
				t.Errorf("getMoneyAmount() = %v, want %v", got, tt.want)
			}
		})
	}
}
