package question_561_570

import (
	"testing"
)

func Test_checkInclusion(t *testing.T) {
	tests := []struct {
		s1   string
		s2   string
		want bool
	}{
		{"ab", "eidbaooo", true},
		{"ab", "eidboaooo", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkInclusion(tt.s1, tt.s2); got != tt.want {
				t.Errorf("checkInclusion() = %v, want %v", got, tt.want)
			}
		})
	}
}
