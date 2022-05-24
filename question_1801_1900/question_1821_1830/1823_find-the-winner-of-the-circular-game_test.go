package question_1821_1830

import (
	"testing"
)

func Test_findTheWinner(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{5, 2, 3},
		{6, 5, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findTheWinner(tt.n, tt.k); got != tt.want {
				t.Errorf("findTheWinner() = %v, want %v", got, tt.want)
			}
		})
	}
}
