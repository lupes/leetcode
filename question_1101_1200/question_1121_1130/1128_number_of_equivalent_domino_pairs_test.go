package question_1121_1130

import (
	"testing"
)

func Test_numEquivDominoPairs(t *testing.T) {
	tests := []struct {
		dominoes [][]int
		want     int
	}{
		{[][]int{{1, 2}, {2, 1}, {3, 4}, {5, 6}}, 1},
		{[][]int{{1, 2}, {2, 1}, {1, 2}, {3, 4}, {5, 6}}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numEquivDominoPairs(tt.dominoes); got != tt.want {
				t.Errorf("numEquivDominoPairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
