package question_911_920

import (
	"testing"
)

func Test_hasGroupsSizeX(t *testing.T) {
	tests := []struct {
		deck []int
		want bool
	}{
		{[]int{1}, false},
		{[]int{1, 2}, false},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 3, 4, 4, 3, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 1}, true},
		{[]int{1, 1, 1, 2, 2, 1, 2, 2, 2, 2}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := hasGroupsSizeX(tt.deck); got != tt.want {
				t.Errorf("hasGroupsSizeX() = %v, want %v", got, tt.want)
			}
		})
	}
}
