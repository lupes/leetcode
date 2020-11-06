package question_1361_1370

import (
	"testing"
)

func Test_rankTeams(t *testing.T) {
	tests := []struct {
		votes []string
		want  string
	}{
		{[]string{"ABC", "ACB", "ABC", "ACB", "ACB"}, "ACB"},
		{[]string{"WXYZ", "XYZW"}, "XWYZ"},
		{[]string{"ZMNAGUEDSJYLBOPHRQICWFXTVK"}, "ZMNAGUEDSJYLBOPHRQICWFXTVK"},
		{[]string{"BCA", "CAB", "CBA", "ABC", "ACB", "BAC"}, "ABC"},
		{[]string{"M", "M", "M", "M"}, "M"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rankTeams(tt.votes); got != tt.want {
				t.Errorf("rankTeams() = %v, want %v", got, tt.want)
			}
		})
	}
}
