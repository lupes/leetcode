package question_141_150

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	got := cache.Get(1)
	if got != 1 {
		t.Errorf("want %d, got:%d \n", 1, got)
	}
	cache.Put(3, 3)
	got = cache.Get(2)
	if got != -1 {
		t.Errorf("want %d, got:%d \n", -1, got)
	}

	cache.Put(4, 4)
	got = cache.Get(1)
	if got != -1 {
		t.Errorf("want %d, got:%d \n", -1, got)
	}

	got = cache.Get(3)
	if got != 3 {
		t.Errorf("want %d, got:%d \n", 3, got)
	}

	got = cache.Get(4)
	if got != 4 {
		t.Errorf("want %d, got:%d \n", 4, got)
	}

}
