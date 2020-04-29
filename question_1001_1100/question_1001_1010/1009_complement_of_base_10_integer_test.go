package question_1001_1010

import (
	"testing"
)

func Test_bitwiseComplement(t *testing.T) {
	tests := []struct {
		N    int
		want int
	}{
		{0, 1},
		{1, 0},
		{5, 2},
		{7, 0},
		{10, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := bitwiseComplement(tt.N); got != tt.want {
				t.Errorf("bitwiseComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
