package question_1181_1190

import (
	"testing"
)

func Test_dayOfTheWeek(t *testing.T) {
	tests := []struct {
		day   int
		month int
		year  int
		want  string
	}{
		{1, 1, 1971, "Friday"},
		{31, 8, 2019, "Saturday"},
		{18, 7, 1999, "Sunday"},
		{15, 8, 1993, "Sunday"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := dayOfTheWeek(tt.day, tt.month, tt.year); got != tt.want {
				t.Errorf("dayOfTheWeek() = %v, want %v", got, tt.want)
			}
		})
	}
}
