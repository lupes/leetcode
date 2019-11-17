package question_491_500

import "testing"

func Test_checkPerfectNumber(t *testing.T) {
	tests := []struct {
		num  int
		want bool
	}{
		{1, false},
		{28, true},
		{27, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkPerfectNumber(tt.num); got != tt.want {
				t.Errorf("checkPerfectNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
