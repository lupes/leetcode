package question_561_570

import (
	"testing"
)

func Test_arrayNesting(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{5, 4, 0, 3, 1, 6, 2}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := arrayNesting(tt.nums); got != tt.want {
				t.Errorf("arrayNesting() = %v, want %v", got, tt.want)
			}
		})
	}
}
