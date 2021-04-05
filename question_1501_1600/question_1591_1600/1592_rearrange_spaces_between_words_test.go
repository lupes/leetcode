package question_1591_1600

import (
	"testing"
)

func Test_reorderSpaces(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{"a", "a"},
		{" a", "a "},
		{"  this   is  a sentence ", "this   is   a   sentence"},
		{" practice   makes   perfect", "practice   makes   perfect "},
		{"hello   world", "hello   world"},
		{"hello   world ", "hello    world"},
		{"  walks  udp package   into  bar a", "walks  udp  package  into  bar  a "},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reorderSpaces(tt.text); got != tt.want {
				t.Errorf("reorderSpaces() = %v, want %v", got, tt.want)
			}
		})
	}
}
