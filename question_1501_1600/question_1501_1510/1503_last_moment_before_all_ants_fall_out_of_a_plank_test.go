package question_1501_1510

import (
	"testing"
)

func Test_getLastMoment(t *testing.T) {
	tests := []struct {
		n     int
		left  []int
		right []int
		want  int
	}{
		{4, []int{4, 3}, []int{0, 1}, 4},
		{7, []int{}, []int{0, 1, 2, 3, 4, 5, 6, 7}, 7},
		{7, []int{}, []int{1, 2, 3, 4, 5, 6, 7}, 6},
		{9, []int{5}, []int{4}, 5},
		{6, []int{6}, []int{0}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getLastMoment(tt.n, tt.left, tt.right); got != tt.want {
				t.Errorf("getLastMoment() = %v, want %v", got, tt.want)
			}
		})
	}
}
