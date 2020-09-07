package question_1261_1270

import (
	"reflect"
	"testing"
)

func Test_suggestedProducts(t *testing.T) {
	tests := []struct {
		products   []string
		searchWord string
		want       [][]string
	}{
		{
			[]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"},
			"mouse",
			[][]string{{"mobile", "moneypot", "monitor"},
				{"mobile", "moneypot", "monitor"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"},
				{"mouse", "mousepad"}},
		},
		{
			[]string{"havana"}, "havana",
			[][]string{{"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}, {"havana"}},
		},
		{
			[]string{"bags", "baggage", "banner", "box", "cloths"},
			"bags",
			[][]string{
				{"baggage", "bags", "banner"},
				{"baggage", "bags", "banner"},
				{"baggage", "bags"}, {"bags"},
			},
		},
		{
			[]string{"havana"}, "tatiana",
			[][]string{{}, {}, {}, {}, {}, {}},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := suggestedProducts(tt.products, tt.searchWord); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("suggestedProducts() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
