package common

import (
	"fmt"
	"runtime"
	"testing"
)

func TestNewBitmap(t *testing.T) {
	got := NewBitmap(1000000000)
	got.Add(0)
	got.Add(1)
	got.Add(63)
	got.Add(64)
	got.Add(65)

	got.Del(1)
	got.Del(1)
	got.Del(9999999999)
	mem := runtime.MemStats{}
	runtime.ReadMemStats(&mem)
	fmt.Printf("%d\n", mem.Alloc)
}
