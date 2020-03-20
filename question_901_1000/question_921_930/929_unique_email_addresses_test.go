package question_921_930

import (
	"testing"
)

func Test_numUniqueEmails(t *testing.T) {
	tests := []struct {
		emails []string
		want   int
	}{
		{[]string{"test.email+alex@leetcode.com", "test.e.mail+bob.cathy@leetcode.com", "testemail+david@lee.tcode.com"}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numUniqueEmails(tt.emails); got != tt.want {
				t.Errorf("numUniqueEmails() = %v, want %v", got, tt.want)
			}
		})
	}
}
