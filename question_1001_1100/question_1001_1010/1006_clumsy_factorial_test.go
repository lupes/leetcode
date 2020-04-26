package question_1001_1010

import (
	"testing"
)

func Test_clumsy(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{1, 1},
		{2, 2},
		{3, 6},
		{4, 7},
		{5, 7},
		{10, 12},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := clumsy(tt.N); got != tt.want {
				t.Errorf("clumsy() = %v, want %v", got, tt.want)
			}
		})
	}
}
