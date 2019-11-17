package question_291_300

import "testing"

func Test_getHint(t *testing.T) {
	tests := []struct {
		name   string
		secret string
		guess  string
		want   string
	}{
		{"test", "1807", "7810", "1A3B"},
		{"test", "1123", "0111", "1A1B"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHint(tt.secret, tt.guess); got != tt.want {
				t.Errorf("getHint() = %v, want %v", got, tt.want)
			}
		})
	}
}
