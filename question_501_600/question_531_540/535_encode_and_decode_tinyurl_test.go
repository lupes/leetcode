package question_531_540

import (
	"fmt"
	"testing"
)

func TestConstructor(t *testing.T) {
	codec := Constructor()
	url := "problems/design-tinyurl"
	key := codec.encode(url)
	fmt.Printf(codec.decode(key))
	if codec.decode(key) != url {
		t.Errorf("want=%+v got=%+v\n", url, codec.decode(key))
	}
}
