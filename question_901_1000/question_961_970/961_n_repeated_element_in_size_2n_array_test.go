package question_961_970

import (
	"testing"
)

func Test_repeatedNTimes(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{1, 2, 3, 3}, 3},
		{[]int{2, 1, 2, 5, 3, 2}, 2},
		{[]int{5, 1, 5, 2, 5, 3, 5, 4}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := repeatedNTimes(tt.A); got != tt.want {
				t.Errorf("repeatedNTimes() = %v, want %v", got, tt.want)
			}
		})
	}
}
