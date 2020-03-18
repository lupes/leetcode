package question_911_920

import (
	"testing"
)

func Test_reverseOnlyLetters(t *testing.T) {
	tests := []struct {
		S    string
		want string
	}{
		{"ab-cd", "dc-ba"},
		{"ab--cd", "dc--ba"},
		{"-ab--cd", "-dc--ba"},
		{"a-bC-dEf-ghIj", "j-Ih-gfE-dCba"},
		{"Test1ng-Leet=code-Q!", "Qedo1ct-eeLg=ntse-T!"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reverseOnlyLetters(tt.S); got != tt.want {
				t.Errorf("reverseOnlyLetters() = %v, want %v", got, tt.want)
			}
		})
	}
}
