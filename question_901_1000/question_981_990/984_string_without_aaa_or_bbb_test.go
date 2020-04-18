package question_981_990

import (
	"testing"
)

func Test_strWithout3a3b(t *testing.T) {
	tests := []struct {
		A    int
		B    int
		want string
	}{
		{1, 2, "bba"},
		{4, 1, "aabaa"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := strWithout3a3b(tt.A, tt.B); got != tt.want {
				t.Errorf("strWithout3a3b() = %v, want %v", got, tt.want)
			}
		})
	}
}
