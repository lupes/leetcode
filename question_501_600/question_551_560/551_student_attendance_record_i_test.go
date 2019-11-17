package question_551_560

import "testing"

func Test_checkRecord(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		{"PPALLP", true},
		{"PPALLL", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkRecord(tt.s); got != tt.want {
				t.Errorf("checkRecord() = %v, want %v", got, tt.want)
			}
		})
	}
}
