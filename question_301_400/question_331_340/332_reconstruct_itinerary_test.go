package question_331_340

import (
	"reflect"
	"testing"
)

func Test_findItinerary(t *testing.T) {
	tests := []struct {
		name    string
		tickets [][]string
		want    []string
	}{
		{"test", [][]string{{"JFK", "KUL"}, {"JFK", "NRT"}, {"NRT", "JFK"}}, []string{"JFK", "NRT", "JFK", "KUL"}},
		{"test", [][]string{{"MUC", "LHR"}, {"JFK", "MUC"}, {"SFO", "SJC"}, {"LHR", "SFO"}}, []string{"JFK", "MUC", "LHR", "SFO", "SJC"}},
		{"test", [][]string{{"JFK", "SFO"}, {"JFK", "ATL"}, {"SFO", "ATL"}, {"ATL", "JFK"}, {"ATL", "SFO"}}, []string{"JFK", "ATL", "JFK", "SFO", "ATL", "SFO"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findItinerary(tt.tickets); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findItinerary() = %v, want %v", got, tt.want)
			}
		})
	}
}
