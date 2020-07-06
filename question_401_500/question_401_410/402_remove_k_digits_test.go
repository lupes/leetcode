package question_401_410

import (
	"testing"
)

func Test_removeKdigits(t *testing.T) {
	tests := []struct {
		num  string
		k    int
		want string
	}{
		{"1432219", 3, "1219"},
		{"1234567890", 9, "0"},
		{"123456789", 8, "1"},
		{"123456789", 9, "0"},
		{"1432219", 4, "119"},
		{"10200", 1, "200"},
		{"10200", 2, "0"},
		{"10", 2, "0"},
		{"100", 1, "0"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeKdigits(tt.num, tt.k); got != tt.want {
				t.Errorf("removeKdigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
