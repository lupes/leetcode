package question_1321_1330

import (
	"testing"
)

func Test_maximum69Number(t *testing.T) {
	tests := []struct {
		num  int
		want int
	}{
		//{9669, 9969},
		{669, 969},
		//{99, 99},
		//{9, 9},
		//{6, 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maximum69Number(tt.num); got != tt.want {
				t.Errorf("maximum69Number() = %v, want %v", got, tt.want)
			}
		})
	}
}
