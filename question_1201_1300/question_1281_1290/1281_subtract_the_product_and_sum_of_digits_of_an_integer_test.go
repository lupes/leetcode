package question_1281_1290

import (
	"testing"
)

func Test_subtractProductAndSum(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{234, 15},
		{4421, 21},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := subtractProductAndSum(tt.n); got != tt.want {
				t.Errorf("subtractProductAndSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
