package question_491_500

import (
	"strconv"
	"testing"
)

func Test_convertToBase7(t *testing.T) {
	tests := []struct {
		name string
		num  int
	}{
		{"test", 0},
		{"test", 1},
		{"test", 10},
		{"test", -7},
		{"test", 100},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			want := strconv.FormatInt(int64(tt.num), 7)

			if got := convertToBase7(tt.num); got != want {
				t.Errorf("convertToBase7() = %v, want %v", got, want)
			} else {
				t.Logf("got = %v, want = %v", got, want)
			}
		})
	}
}
