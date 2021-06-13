package question_371_380

import (
	"testing"
)

func Test_getSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 2, 1 + 2},
		{2, -2, 2 - 2},
		{-2, 3, 3 - 2},
		{123456, 234567, 358023},
		{123491856, 234711567, 123491856 + 234711567},
		{123491856, -234711567, 123491856 - 234711567},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getSum(tt.a, tt.b); got != tt.want {
				t.Errorf("getSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
