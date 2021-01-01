package question_1211_1220

import (
	"testing"
)

func Test_minCostToMoveChips(t *testing.T) {
	tests := []struct {
		chips []int
		want  int
	}{
		{[]int{1, 2, 3}, 1},
		{[]int{2, 2, 2, 3, 3}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minCostToMoveChips(tt.chips); got != tt.want {
				t.Errorf("minCostToMoveChips() = %v, want %v", got, tt.want)
			}
		})
	}
}
