package question_841_850

import (
	"testing"
)

func Test_maxDistToClosest(t *testing.T) {
	tests := []struct {
		seats []int
		want  int
	}{
		{[]int{1, 0, 0, 0, 1, 0, 1}, 2},
		{[]int{1, 0, 0, 0}, 3},
		{[]int{1, 0, 1}, 1},
		{[]int{0, 0, 1, 0, 1}, 2},
		{[]int{0, 1, 0, 1, 0, 0, 0, 0, 0, 1}, 3},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxDistToClosest(tt.seats); got != tt.want {
				t.Errorf("maxDistToClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
