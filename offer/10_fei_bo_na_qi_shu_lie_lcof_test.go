package offer

import (
	"testing"
)

func Test_fib(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{43, 433494437},
		{45, 134903163},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := fib(tt.n); got != tt.want {
				t.Errorf("fib() = %v, want %v", got, tt.want)
			}
		})
	}
}
