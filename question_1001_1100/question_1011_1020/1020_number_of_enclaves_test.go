package question_1011_1020

import (
	"testing"
)

func Test_numEnclaves(t *testing.T) {
	tests := []struct {
		A    [][]int
		want int
	}{
		{[][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 1, 1, 0}, {0, 0, 0, 0}}, 3},
		{[][]int{{0, 1, 1, 0}, {0, 0, 1, 0}, {0, 0, 1, 0}, {0, 0, 0, 0}}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numEnclaves(tt.A); got != tt.want {
				t.Errorf("numEnclaves() = %v, want %v", got, tt.want)
			}
		})
	}
}
