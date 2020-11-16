package question_1331_1340

import (
	"testing"
)

func Test_isPossible(t *testing.T) {
	tests := []struct {
		target []int
		want   bool
	}{
		{[]int{9, 3, 5}, true},
		{[]int{1, 1, 1, 2}, false},
		{[]int{1, 1, 2}, false},
		{[]int{5, 8}, true},
		{[]int{1, 1000000000}, true},
		{[]int{2, 900000001}, true},
		{[]int{1}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isPossible(tt.target); got != tt.want {
				t.Errorf("isPossible() = %v, want %v", got, tt.want)
			}
		})
	}
}
