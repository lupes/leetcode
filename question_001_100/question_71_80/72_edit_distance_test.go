package question_71_80

import "testing"

func Test_minDistance(t *testing.T) {
	tests := []struct {
		word1 string
		word2 string
		want  int
	}{
		{"a", "ab", 1},
		{"horse", "ros", 3},
		{"intention", "execution", 5},
		{"intentione", "execution", 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDistance(tt.word1, tt.word2); got != tt.want {
				t.Errorf("minDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
