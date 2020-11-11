package question_1331_1340

import (
	"testing"
)

func Test_daysBetweenDates(t *testing.T) {
	tests := []struct {
		date1 string
		date2 string
		want  int
	}{
		{"2019-06-29", "2019-06-30", 1},
		{"2020-01-15", "2019-12-31", 15},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := daysBetweenDates(tt.date1, tt.date2); got != tt.want {
				t.Errorf("daysBetweenDates() = %v, want %v", got, tt.want)
			}
		})
	}
}
