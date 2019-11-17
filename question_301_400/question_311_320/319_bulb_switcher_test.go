package question_311_320

import "testing"

func Test_bulbSwitch(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"test", 3, 1},
		{"test", 4, 2},
		{"test", 5, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := bulbSwitch(tt.n); got != tt.want {
				t.Errorf("bulbSwitch() = %v, want %v", got, tt.want)
			}
		})
	}
}
