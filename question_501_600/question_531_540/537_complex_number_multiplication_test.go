package question_531_540

import (
	"testing"
)

func Test_toInt(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"-12972", -12972},
		{"12972", 12972},
		{"0", 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := toInt(tt.s); got != tt.want {
				t.Errorf("toInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseComplexNum(t *testing.T) {
	tests := []struct {
		s     string
		wantA int
		wantB int
	}{
		{"1+1i", 1, 1},
		{"1", 1, 0},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			gotA, gotB := parseComplexNum(tt.s)
			if gotA != tt.wantA {
				t.Errorf("parseComplexNum() gotA = %v, want %v", gotA, tt.wantA)
			}
			if gotB != tt.wantB {
				t.Errorf("parseComplexNum() gotB = %v, want %v", gotB, tt.wantB)
			}
		})
	}
}

func Test_toStr(t *testing.T) {
	tests := []struct {
		i    int
		want string
	}{
		{123, "123"},
		{-123, "-123"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := toStr(tt.i); got != tt.want {
				t.Errorf("toStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_complexNumberMultiply(t *testing.T) {
	tests := []struct {
		a    string
		b    string
		want string
	}{
		{"1+1i", "1+1i", "0+2i"},
		{"1+-1i", "1+-1i", "0+-2i"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := complexNumberMultiply(tt.a, tt.b); got != tt.want {
				t.Errorf("complexNumberMultiply() = %v, want %v", got, tt.want)
			}
		})
	}
}
