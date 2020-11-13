package question_1331_1340

import (
	"testing"
)

func Test_countOrders(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 6},
		{3, 90},
		{100, 14159051},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countOrders(tt.n); got != tt.want {
				t.Errorf("countOrders() = %v, want %v", got, tt.want)
			}
		})
	}
}
