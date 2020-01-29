package question_741_750

import (
	"testing"
)

func Test_nextGreatestLetter(t *testing.T) {
	tests := []struct {
		letters []byte
		target  byte
		want    byte
	}{
		{[]byte{'b', 'd'}, 'a', 'b'},
		{[]byte{'b', 'b', 'b', 'b', 'd'}, 'b', 'd'},
		{[]byte{'b', 'd'}, 'a', 'b'},
		{[]byte{'c', 'f', 'j'}, 'a', 'c'},
		{[]byte{'c', 'f', 'j'}, 'c', 'f'},
		{[]byte{'c', 'f', 'j'}, 'd', 'f'},
		{[]byte{'c', 'f', 'j'}, 'g', 'j'},
		{[]byte{'c', 'f', 'j'}, 'j', 'c'},
		{[]byte{'c', 'f', 'j'}, 'k', 'c'},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := nextGreatestLetter(tt.letters, tt.target); got != tt.want {
				t.Errorf("nextGreatestLetter() = %v, want %v", got, tt.want)
			}
		})
	}
}
