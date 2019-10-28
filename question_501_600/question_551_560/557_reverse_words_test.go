package question_551_560

import "testing"

func Test_reverseWords(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"Let's take LeetCode contest", "s'teL ekat edoCteeL tsetnoc"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reverseWords(tt.s); got != tt.want {
				t.Errorf("reverseWords() = [%v], want [%v]", got, tt.want)
			}
		})
	}
}
