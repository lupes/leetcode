package question_291_300

import "testing"

func Test_canWinNim(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name string
		n    int
		want bool
	}{
		{"test", 1, true},
		{"test", 2, true},
		{"test", 3, true},
		{"test", 4, false},
		{"test", 5, true},
		{"test", 10, true},
		{"test", 15, true},
		{"test", 16, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canWinNim(tt.n); got != tt.want {
				t.Errorf("canWinNim() = %v, want %v", got, tt.want)
			}
		})
	}
}
