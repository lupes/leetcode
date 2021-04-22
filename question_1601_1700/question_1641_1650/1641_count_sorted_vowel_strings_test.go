package question_1641_1650

import (
	"testing"
)

func Test_countVowelStrings(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{1, 5},
		{2, 15},
		{33, 66045},
		{50, 316251},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countVowelStrings(tt.n); got != tt.want {
				t.Errorf("countVowelStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}
