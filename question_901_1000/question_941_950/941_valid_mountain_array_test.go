package question_941_950

import (
	"testing"
)

func Test_validMountainArray(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		{[]int{1, 2, 1}, true},
		{[]int{2, 1}, false},
		{[]int{1, 2, 3}, false},
		{[]int{3, 5, 5}, false},
		{[]int{1, 2, 3, 2}, true},
		{[]int{0, 3, 2, 1}, true},
		{[]int{14, 82, 89, 84, 79, 70, 70, 68, 67, 66, 63, 60, 58, 54, 44, 43, 32, 28, 26, 25, 22, 15, 13, 12, 10, 8, 7, 5, 4, 3}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := validMountainArray(tt.A); got != tt.want {
				t.Errorf("validMountainArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
