package question_861_870

import (
	"testing"
)

func Test_reorderedPowerOf2(t *testing.T) {
	tests := []struct {
		N    int
		want bool
	}{
		{1, true},
		{46, true},
		{46002826, false},
		{271612776, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reorderedPowerOf2(tt.N); got != tt.want {
				t.Errorf("reorderedPowerOf2() = %v, want %v", got, tt.want)
			}
		})
	}
}
