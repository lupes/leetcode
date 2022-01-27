package question_2041_2050

import (
	"testing"
)

func Test_countValidWords(t *testing.T) {
	tests := []struct {
		sentence string
		want     int
	}{
		{"!", 1},
		{"cat and  dog", 3},
		{"!this  1-s b8d!", 0},
		{"alice and  bob are playing stone-game10", 5},
		{"he bought 2 pencils, 3 erasers, and 1  pencil-sharpener.", 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := countValidWords(tt.sentence); got != tt.want {
				t.Errorf("countValidWords() = %v, want %v", got, tt.want)
			}
		})
	}
}
