package question_2021_2030

import (
	"testing"
)

func Test_maxConsecutiveAnswers(t *testing.T) {
	tests := []struct {
		answerKey string
		k         int
		want      int
	}{
		{"TTFF", 2, 4},
		{"T", 1, 1},
		{"TFFT", 1, 3},
		{"TTFTTFTT", 1, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxConsecutiveAnswers(tt.answerKey, tt.k); got != tt.want {
				t.Errorf("maxConsecutiveAnswers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_maxConsecutiveAnswersHelper(t *testing.T) {
	tests := []struct {
		answerKey string
		k         int
		b         int32
		want      int
	}{
		{"TTFF", 2, 'T', 4},
		{"TTFF", 1, 'T', 3},
		{"TFFT", 1, 'F', 2},
		{"TTFTTFTT", 1, 'T', 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxConsecutiveAnswersHelper(tt.answerKey, tt.k, tt.b); got != tt.want {
				t.Errorf("maxConsecutiveAnswersHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}
