package question_421_430

import (
	"testing"
)

func Test_originalDigits(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"owoztneoer", "012"},
		{"fviefuro", "45"},
		{"zeroonetwothreefourfivesixseveneightnine", "0123456789"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := originalDigits(tt.s); got != tt.want {
				t.Errorf("originalDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
