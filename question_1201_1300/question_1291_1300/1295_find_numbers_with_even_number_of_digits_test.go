package question_1291_1300

import (
	"testing"
)

func Test_findNumbers(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findNumbers(tt.nums); got != tt.want {
				t.Errorf("findNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
