package question_811_820

import (
	"reflect"
	"sort"
	"testing"
)

func Test_subdomainVisits(t *testing.T) {
	tests := []struct {
		cpdomains []string
		want      []string
	}{
		{[]string{"9001 discuss.leetcode.com"}, []string{"9001 discuss.leetcode.com", "9001 leetcode.com", "9001 com"}},
		{[]string{"900 google.mail.com", "50 yahoo.com", "1 intel.mail.com", "5 wiki.org"}, []string{"901 mail.com", "50 yahoo.com", "900 google.mail.com", "5 wiki.org", "5 org", "1 intel.mail.com", "951 com"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := subdomainVisits(tt.cpdomains)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("subdomainVisits() = %v, want %v", got, tt.want)
			}
		})
	}
}
