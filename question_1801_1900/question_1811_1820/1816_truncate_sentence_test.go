package question_1811_1820

import (
	"testing"
)

func Test_truncateSentence(t *testing.T) {
	tests := []struct {
		s    string
		k    int
		want string
	}{
		{"Hello how are you Contestant", 4, "Hello how are you"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := truncateSentence(tt.s, tt.k); got != tt.want {
				t.Errorf("truncateSentence() = %v, want %v", got, tt.want)
			}
		})
	}
}
