package question_1171_1180

import (
	"testing"
)

func Test_numPrimeArrangements(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{5, 12},
		{6, 36},
		{100, 682289015},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numPrimeArrangements(tt.n); got != tt.want {
				t.Errorf("numPrimeArrangements() = %v, want %v", got, tt.want)
			}
		})
	}
}
