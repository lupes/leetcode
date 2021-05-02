package question_1661_1670

import (
	"testing"
)

func Test_minimumEffort(t *testing.T) {
	tests := []struct {
		tasks [][]int
		want  int
	}{
		{[][]int{{1, 2}, {2, 4}, {4, 8}, {3, 8}}, 11},
		{[][]int{{1, 2}, {2, 4}, {4, 8}}, 8},
		{[][]int{{1, 3}, {2, 4}, {10, 11}, {10, 12}, {8, 9}}, 32},
		{[][]int{{1, 7}, {2, 8}, {3, 9}, {4, 10}, {5, 11}, {6, 12}}, 27},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minimumEffort(tt.tasks); got != tt.want {
				t.Errorf("minimumEffort() = %v, want %v", got, tt.want)
			}
		})
	}
}
