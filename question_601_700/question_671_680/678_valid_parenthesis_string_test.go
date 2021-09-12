package question_671_680

import (
	"testing"
)

func Test_checkValidString(t *testing.T) {
	tests := []struct {
		s    string
		want bool
	}{
		//{"()", true},
		//{"(*)", true},
		//{"(*))", true},
		{"(((((()*)(*)*))())())(()())())))((**)))))(()())()", true},
		{"((***)))))(**))))", true},
		//{"(((((*(()((((*((**(((()()*)()()()*((((**)())*)*)))))))(())(()))())((*()()(((()((()*(())*(()**)()(())", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkValidString(tt.s); got != tt.want {
				t.Errorf("checkValidString() = %v, want %v", got, tt.want)
			}
		})
	}
}
