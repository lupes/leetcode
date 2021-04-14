package common

import (
	"fmt"
	"math/rand"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func genNNums(n int, max int) []int {
	var nums = make([]int, 0, n)
	for i := 0; i < n; i++ {
		nums = append(nums, rand.Intn(max))
	}
	return nums
}

func Test_sort(t *testing.T) {
	for _, sort := range []Sort{bobbleSort, insertionSort, selectionSort, mergeSort, quickSort, countingSort} {
		var times time.Duration
		for i := 0; i < 1; i++ {
			nums := genNNums(100000, 10000)
			start := time.Now()
			sort(nums)
			times += time.Now().Sub(start)
		}
		fmt.Printf("%+v:%s\n", runtime.FuncForPC(reflect.ValueOf(sort).Pointer()).Name(), times/10)
	}
}
