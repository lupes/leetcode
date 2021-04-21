package question_1621_1630

import (
	"testing"
)

func Test_slowestKey(t *testing.T) {
	tests := []struct {
		releaseTimes []int
		keysPressed  string
		want         byte
	}{
		{[]int{9, 29, 49, 50}, "cbcd", 'c'},
		{[]int{12, 23, 36, 46, 62}, "spuda", 'a'},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := slowestKey(tt.releaseTimes, tt.keysPressed); got != tt.want {
				t.Errorf("slowestKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
