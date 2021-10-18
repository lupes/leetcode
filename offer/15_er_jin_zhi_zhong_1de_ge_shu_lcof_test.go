package offer

import (
	"testing"
)

func Test_hammingWeight(t *testing.T) {
	tests := []struct {
		num  uint32
		want int
	}{
		{11, 3},
		{128, 1},
		{4294967293, 31},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := hammingWeight(tt.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}
