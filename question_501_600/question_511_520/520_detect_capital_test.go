package qustion_511_520

import "testing"

func Test_detectCapitalUse(t *testing.T) {
	tests := []struct {
		word string
		want bool
	}{
		{"test", true},
		{"USA", true},
		{"Google", true},
		{"GooGle", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := detectCapitalUse(tt.word); got != tt.want {
				t.Errorf("detectCapitalUse() = %v, want %v", got, tt.want)
			}
		})
	}
}
