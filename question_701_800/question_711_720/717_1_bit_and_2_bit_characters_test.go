package question_711_720

import (
	"testing"
)

func Test_isOneBitCharacter(t *testing.T) {
	tests := []struct {
		bits []int
		want bool
	}{
		{[]int{0}, true},
		{[]int{0, 0}, true},
		{[]int{1, 0, 0}, true},
		{[]int{1, 0, 1, 0}, false},
		{[]int{1, 1, 1, 0}, false},
		{[]int{1, 1, 1, 0, 0}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isOneBitCharacter(tt.bits); got != tt.want {
				t.Errorf("isOneBitCharacter() = %v, want %v", got, tt.want)
			}
		})
	}
}
