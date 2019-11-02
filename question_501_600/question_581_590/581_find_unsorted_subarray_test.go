package question_581_590

import "testing"

func Test_findUnsortedSubarray(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1}, 0},
		{[]int{1, 2}, 0},
		{[]int{1, 2, 3, 5, 4}, 2},
		{[]int{1, 2, 3, 3, 3}, 0},
		{[]int{2, 6, 4, 8, 10, 9, 15}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findUnsortedSubarray(tt.nums); got != tt.want {
				t.Errorf("findUnsortedSubarray() = %v, want %v", got, tt.want)
			}
		})
	}
}
