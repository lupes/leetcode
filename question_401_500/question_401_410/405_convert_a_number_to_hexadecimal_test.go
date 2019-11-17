package question_401_410

import "testing"

func Test_toHex(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want string
	}{
		{"test", 26, "1a"},
		{"test", -1, "ffffffff"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toHex(tt.num); got != tt.want {
				t.Errorf("toHex() = %v, want %v", got, tt.want)
			}
		})
	}
}
