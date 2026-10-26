package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"sort"
	"testing"
)

var checks int

func eq(t *testing.T, got, want any) {
	t.Helper()
	checks++
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v (check %d)", got, want, checks)
	}
}
func forEachArray(maxLength int, alphabet []int, visit func([]int)) {
	var walk func([]int)
	walk = func(a []int) {
		visit(a)
		if len(a) == maxLength {
			return
		}
		for _, value := range alphabet {
			next := append(append([]int{}, a...), value)
			walk(next)
		}
	}
	walk([]int{})
}
func TestExamplesAndBoundaries(t *testing.T) {
	defer testutil.SkipTODO(t)
	eq(t, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	eq(t, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
	forEachArray(6, []int{-2, 0, 1, 2}, func(a []int) {
		if len(a) == 0 {
			return
		}
		ordered := append([]int{}, a...)
		sort.Ints(ordered)
		for k := 1; k <= len(a); k++ {
			input := append([]int{}, a...)
			eq(t, findKthLargest(input, k), ordered[len(a)-k])
		}
	})
	eq(t, findKthLargest([]int{-10000, 10000, 10000}, 2), 10000)
	t.Logf("checks=%d", checks)
}
