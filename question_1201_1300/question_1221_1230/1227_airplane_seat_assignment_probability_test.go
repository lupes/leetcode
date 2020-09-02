package question_1221_1230

import (
	"testing"
)

func Test_nthPersonGetsNthSeat(t *testing.T) {
	tests := []struct {
		n    int
		want float64
	}{
		{1, 1},
		{2, 0.5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := nthPersonGetsNthSeat(tt.n); got != tt.want {
				t.Errorf("nthPersonGetsNthSeat() = %v, want %v", got, tt.want)
			}
		})
	}
}
