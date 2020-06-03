package question_1161_1170

import (
	"reflect"
	"sort"
	"testing"
)

func Test_invalidTransactions(t *testing.T) {
	tests := []struct {
		transactions []string
		want         []string
	}{
		{[]string{"alice,20,800,mtv", "alice,50,100,beijing"}, []string{"alice,20,800,mtv", "alice,50,100,beijing"}},
		{[]string{"alice,20,800,mtv", "alice,50,1200,mtv"}, []string{"alice,50,1200,mtv"}},
		{[]string{"alice,20,800,mtv", "bob,50,1200,mtv"}, []string{"bob,50,1200,mtv"}},
		{[]string{"alex,676,260,bangkok", "bob,656,1366,bangkok", "alex,393,616,bangkok",
			"bob,820,990,amsterdam", "alex,596,1390,amsterdam"}, []string{"bob,50,1200,mtv"}},
		{[]string{"bob,689,1910,barcelona", "alex,696,122,bangkok", "bob,832,1726,barcelona",
			"bob,820,596,bangkok", "chalicefy,217,669,barcelona", "bob,175,221,amsterdam"},
			[]string{"bob,689,1910,barcelona", "bob,832,1726,barcelona", "bob,820,596,bangkok"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := invalidTransactions(tt.transactions)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("invalidTransactions() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
