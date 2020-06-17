package question_1251_1260

import (
	"testing"
)

func Test_oddCells(t *testing.T) {
	tests := []struct {
		n       int
		m       int
		indices [][]int
		want    int
	}{
		{2, 3, [][]int{{0, 1}, {1, 1}}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := oddCells(tt.n, tt.m, tt.indices); got != tt.want {
				t.Errorf("oddCells() = %v, want %v", got, tt.want)
			}
		})
	}
}
