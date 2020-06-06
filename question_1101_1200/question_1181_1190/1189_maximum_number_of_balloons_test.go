package question_1181_1190

import (
	"testing"
)

func Test_maxNumberOfBalloons(t *testing.T) {
	tests := []struct {
		text string
		want int
	}{
		{"nlaebolko", 1},
		{"loonbalxballpoon", 2},
		{"leetcode", 0},
		{"balon", 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxNumberOfBalloons(tt.text); got != tt.want {
				t.Errorf("maxNumberOfBalloons() = %v, want %v", got, tt.want)
			}
		})
	}
}
