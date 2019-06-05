package question_91_100

import "testing"

func Test_numDecodings(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test", "0", 0},
		{"test", "01", 0},
		{"test", "00", 0},
		{"test", "000", 0},
		{"test", "110", 1},
		{"test", "611", 2},
		{"test", "1000", 0},
		{"test", "12", 2},
		{"test", "10", 1},
		{"test", "27", 1},
		{"test", "100", 0},
		{"test", "1010", 1},
		{"test", "226", 3},
		{"test", "2260", 0},
		{"test", "2260", 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numDecodings(tt.s); got != tt.want {
				t.Errorf("numDecodings() = %v, want %v", got, tt.want)
			}
		})
	}
}
