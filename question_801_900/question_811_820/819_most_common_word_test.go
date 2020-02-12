package question_811_820

import (
	"testing"
)

func Test_mostCommonWord(t *testing.T) {
	tests := []struct {
		paragraph string
		banned    []string
		want      string
	}{
		{"Bob hit a ball, the hit BALL flew far after it was hit. her's", []string{"hit"}, "ball"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := mostCommonWord(tt.paragraph, tt.banned); got != tt.want {
				t.Errorf("mostCommonWord() = %v, want %v", got, tt.want)
			}
		})
	}
}
