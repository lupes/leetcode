package question_1501_1510

import (
	"testing"
)

func Test_reformatDate(t *testing.T) {
	tests := []struct {
		date string
		want string
	}{
		{"20th Oct 2052", "2052-10-20"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reformatDate(tt.date); got != tt.want {
				t.Errorf("reformatDate() = %v, want %v", got, tt.want)
			}
		})
	}
}
