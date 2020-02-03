package question_771_780

import (
	"testing"
)

func Test_numJewelsInStones(t *testing.T) {
	tests := []struct {
		J    string
		S    string
		want int
	}{
		{"aA", "aAAbbbb", 3},
		{"z", "ZZ", 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numJewelsInStones(tt.J, tt.S); got != tt.want {
				t.Errorf("numJewelsInStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
