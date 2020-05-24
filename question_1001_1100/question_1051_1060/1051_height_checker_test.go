package question_1051_1060

import (
	"testing"
)

func Test_heightChecker(t *testing.T) {
	tests := []struct {
		heights []int
		want    int
	}{
		{[]int{1, 1, 4, 2, 1, 3}, 3},
		{[]int{1, 2, 3, 4, 5}, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := heightChecker(tt.heights); got != tt.want {
				t.Errorf("heightChecker() = %v, want %v", got, tt.want)
			}
		})
	}
}
