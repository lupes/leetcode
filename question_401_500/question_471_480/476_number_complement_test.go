package question_471_480

import (
	"testing"
)

func Test_findComplement(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{5, 2},
		{1, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findComplement(tt.num); got != tt.want {
				t.Errorf("findComplement() = %v, want %v", got, tt.want)
			}
		})
	}
}
