package question_1421_1430

import (
	"testing"
)

func Test_maxScore(t *testing.T) {
	tests := []struct {
		cardPoints []int
		k          int
		want       int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 1}, 3, 12},
		{[]int{2, 2, 2}, 1, 2},
		{[]int{9, 7, 7, 9, 7, 7, 9}, 7, 55},
		{[]int{1, 79, 80, 1, 1, 1, 200, 1}, 3, 202},
		{[]int{1, 79, 80, 80, 1, 1, 200, 1}, 4, 281},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxScore(tt.cardPoints, tt.k); got != tt.want {
				t.Errorf("maxScore() = %v, want %v", got, tt.want)
			}
		})
	}
}
