package question_1661_1670

import (
	"testing"
)

func Test_arrayStringsAreEqual(t *testing.T) {
	tests := []struct {
		word1 []string
		word2 []string
		want  bool
	}{
		{[]string{"ab", "c"}, []string{"a", "bc"}, true},
		{[]string{"a", "cb"}, []string{"ab", "c"}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := arrayStringsAreEqual(tt.word1, tt.word2); got != tt.want {
				t.Errorf("arrayStringsAreEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
