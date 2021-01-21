package question_1401_1410

import (
	"testing"
)

func Test_entityParser(t *testing.T) {
	tests := []struct {
		text string
		want string
	}{
		{"&amp;gt;", "&gt;"},
		{"&amp; is an HTML entity but &ambassador; is not.", `& is an HTML entity but &ambassador; is not.`},
		{"and I quote: &quot;...&quot;", `and I quote: "..."`},
		{"Stay home! Practice on Leetcode :)", `Stay home! Practice on Leetcode :)`},
		{"leetcode.com&frasl;problemset&frasl;all", `leetcode.com/problemset/all`},
		{"x &gt; y &amp;&amp; x &lt; y is always false", `x > y && x < y is always false`},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := entityParser(tt.text); got != tt.want {
				t.Errorf("entityParser() = %v, want %v", got, tt.want)
			}
		})
	}
}
