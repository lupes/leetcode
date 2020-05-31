package question_1151_1160

import (
	"testing"
)

func Test_dayOfYear(t *testing.T) {
	tests := []struct {
		date string
		want int
	}{
		{"2019-01-09", 9},
		{"2019-02-10", 41},
		{"2003-03-01", 60},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := dayOfYear(tt.date); got != tt.want {
				t.Errorf("dayOfYear() = %v, want %v", got, tt.want)
			}
		})
	}
}
