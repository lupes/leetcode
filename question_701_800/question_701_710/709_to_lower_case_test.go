package question_701_710

import (
	"testing"
)

func Test_toLowerCase(t *testing.T) {
	tests := []struct {
		str  string
		want string
	}{
		{"Hello", "hello"},
		{"here", "here"},
		{"LOVELY", "lovely"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := toLowerCase(tt.str); got != tt.want {
				t.Errorf("toLowerCase() = %v, want %v", got, tt.want)
			}
		})
	}
}
