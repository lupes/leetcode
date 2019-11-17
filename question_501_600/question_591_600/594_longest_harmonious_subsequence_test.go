package question_591_600

import "testing"

func Test_findLHS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{1, 3, 2, 2, 5, 2, 3, 7}, 5},
		{[]int{1, 3, 2, 2, 5, 2, 4, 7}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findLHS(tt.nums); got != tt.want {
				t.Errorf("findLHS() = %v, want %v", got, tt.want)
			}
		})
	}
}
