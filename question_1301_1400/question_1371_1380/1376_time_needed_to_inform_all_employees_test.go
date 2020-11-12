package question_1361_1370

import (
	"testing"
)

func Test_numOfMinutes(t *testing.T) {
	tests := []struct {
		n          int
		headID     int
		manager    []int
		informTime []int
		want       int
	}{
		{1, 0, []int{-1}, []int{0}, 0},
		{6, 2, []int{2, 2, -1, 2, 2, 2}, []int{0, 0, 1, 0, 0, 0}, 1},
		{6, 2, []int{1, 2, 3, 4, 5, 6, -1}, []int{0, 6, 5, 4, 3, 2, 1}, 21},
		{6, 2, []int{-1, 0, 0, 1, 1, 2, 2, 3, 3, 4, 4, 5, 5, 6, 6}, []int{1, 1, 1, 1, 1, 1, 1, 0, 0, 0, 0, 0, 0, 0, 0}, 3},
		{6, 2, []int{3, 3, -1, 2}, []int{0, 0, 162, 914}, 1076},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numOfMinutes(tt.n, tt.headID, tt.manager, tt.informTime); got != tt.want {
				t.Errorf("numOfMinutes() = %v, want %v", got, tt.want)
			}
		})
	}
}
