package question_1211_1220

import (
	"testing"
)

func Test_longestSubsequence(t *testing.T) {
	tests := []struct {
		arr        []int
		difference int
		want       int
	}{
		{[]int{1, 2, 3, 4}, 1, 4},
		{[]int{1, 3, 5, 7}, 1, 1},
		{[]int{1, 5, 7, 8, 5, 3, 4, 2, 1}, -2, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := longestSubsequence(tt.arr, tt.difference); got != tt.want {
				t.Errorf("longestSubsequence() = %v, want %v", got, tt.want)
			}
		})
	}
}
