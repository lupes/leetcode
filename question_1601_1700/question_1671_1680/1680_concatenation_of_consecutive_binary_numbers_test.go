package question_1671_1680

import (
	"testing"
)

func Test_concatenatedBinary(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 1},
		{2, 6},
		{3, 27},
		{4, 220},
		{12, 505379714},
		{100000, 757631812},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := concatenatedBinary(tt.n); got != tt.want {
				t.Errorf("concatenatedBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}
