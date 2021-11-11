package question_621_630

import (
	"testing"
)

func Test_kInversePairs(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{3, 0, 1},
		{3, 1, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kInversePairs(tt.n, tt.k); got != tt.want {
				t.Errorf("kInversePairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
