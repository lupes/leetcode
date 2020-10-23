package question_1331_1340

import (
	"testing"
)

func Test_numberOfSteps(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		{14, 6},
		{8, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfSteps(tt.num); got != tt.want {
				t.Errorf("numberOfSteps() = %v, want %v", got, tt.want)
			}
		})
	}
}
