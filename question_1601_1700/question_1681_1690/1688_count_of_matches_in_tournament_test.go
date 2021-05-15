package question_1681_1690

import (
	"testing"
)

func Test_numberOfMatches(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{7, 6},
		{14, 13},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfMatches(tt.n); got != tt.want {
				t.Errorf("numberOfMatches() = %v, want %v", got, tt.want)
			}
		})
	}
}
