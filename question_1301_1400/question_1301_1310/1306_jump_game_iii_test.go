package question_1301_1310

import (
	"testing"
)

func Test_canReach(t *testing.T) {
	tests := []struct {
		arr   []int
		start int
		want  bool
	}{
		{[]int{4, 2, 3, 0, 3, 1, 2}, 5, true},
		{[]int{4, 2, 3, 0, 3, 1, 2}, 0, true},
		{[]int{3, 0, 2, 1, 2}, 2, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canReach(tt.arr, tt.start); got != tt.want {
				t.Errorf("canReach() = %v, want %v", got, tt.want)
			}
		})
	}
}
