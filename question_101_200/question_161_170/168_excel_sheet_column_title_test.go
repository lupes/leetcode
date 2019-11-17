package question_161_170

import "testing"

func Test_convertToTitle(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want string
	}{
		{"test", 1, "A"},
		{"test", 2, "B"},
		{"test", 3, "C"},
		{"test", 26, "Z"},
		{"test", 27, "AA"},
		{"test", 28, "AB"},
		{"test", 52, "AZ"},
		{"test", 701, "ZY"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := convertToTitle(tt.n); got != tt.want {
				t.Errorf("convertToTitle() = %v, want %v", got, tt.want)
			}
		})
	}
}
