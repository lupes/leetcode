package question_941_950

import (
	"testing"
)

func Test_largestTimeFromDigits(t *testing.T) {
	tests := []struct {
		A    []int
		want string
	}{
		{[]int{1, 2, 3, 4}, "23:41"},
		{[]int{3, 4, 5, 6}, ""},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := largestTimeFromDigits(tt.A); got != tt.want {
				t.Errorf("largestTimeFromDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
