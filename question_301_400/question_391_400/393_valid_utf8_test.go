package question_391_400

import "testing"

func Test_validUtf8(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want bool
	}{
		{"test", []int{0x11}, true},
		{"test", []int{0xA1}, false},
		{"test", []int{197, 130, 1}, true},
		{"test", []int{235, 140, 4}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := validUtf8(tt.data); got != tt.want {
				t.Errorf("validUtf8() = %v, want %v", got, tt.want)
			}
		})
	}
}
