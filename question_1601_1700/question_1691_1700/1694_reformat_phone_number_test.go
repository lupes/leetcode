package question_1691_1700

import (
	"testing"
)

func Test_reformatNumber(t *testing.T) {
	tests := []struct {
		number string
		want   string
	}{
		{"1-23-45 6", "123-456"},
		{"123 4-567", "123-45-67"},
		{"123 4-5678", "123-456-78"},
		{"12", "12"},
		{"--17-5 229 35-39475 ", "175-229-353-94-75"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reformatNumber(tt.number); got != tt.want {
				t.Errorf("reformatNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
