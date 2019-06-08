package question_91_100

import (
	"reflect"
	"testing"
)

func Test_restoreIpAddresses(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want []string
	}{
		{"test", "0000", []string{"0.0.0.0"}},
		{"test", "010010", []string{"0.10.0.10", "0.100.1.0"}},
		{"test", "0000", []string{"0.0.0.0"}},
		{"test", "11111", []string{"1.1.1.11", "1.1.11.1", "1.11.1.1", "11.1.1.1"}},
		{"test", "25525511135", []string{"255.255.11.135", "255.255.111.35"}},
		{"test", "255255111135", []string{"255.255.111.135"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := restoreIpAddresses(tt.s); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("restoreIpAddresses() = %#v, want %#v", got, tt.want)
			}
		})
	}
}

func Test_validIpAddresses(t *testing.T) {
	tests := []struct {
		name  string
		s     string
		i     int
		j     int
		k     int
		want  string
		want1 bool
	}{
		{"test", "1111", 1, 2, 3, "1.1.1.1", true},
		{"test", "25525511135", 1, 2, 3, "", false},
		{"test", "25525511135", 1, 2, 4, "", false},
		{"test", "25525511135", 3, 6, 9, "255.255.111.35", true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := validIpAddresses(tt.s, tt.i, tt.j, tt.k)
			if got != tt.want {
				t.Errorf("validIpAddresses() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("validIpAddresses() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
