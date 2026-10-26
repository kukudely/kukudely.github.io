package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
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
	eq(t, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}), []int{1, 1, 4, 2, 1, 1, 0, 0})
	eq(t, dailyTemperatures([]int{30, 40, 50, 60}), []int{1, 1, 1, 0})
	forEachArray(7, []int{30, 31, 32}, func(a []int) {
		expected := make([]int, len(a))
		for i := range a {
			for j := i + 1; j < len(a); j++ {
				if a[j] > a[i] {
					expected[i] = j - i
					break
				}
			}
		}
		eq(t, dailyTemperatures(a), expected)
	})
	t.Logf("checks=%d", checks)
}
