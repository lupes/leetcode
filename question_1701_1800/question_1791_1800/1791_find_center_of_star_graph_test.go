package question_1791_1800

import (
	"testing"
)

func Test_findCenter(t *testing.T) {
	tests := []struct {
		edges [][]int
		want  int
	}{
		{[][]int{{1, 2}, {2, 3}, {4, 2}}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findCenter(tt.edges); got != tt.want {
				t.Errorf("findCenter() = %v, want %v", got, tt.want)
			}
		})
	}
}
