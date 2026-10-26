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
	eq(t, largestRectangleArea([]int{2, 1, 5, 6, 2, 3}), 10)
	eq(t, largestRectangleArea([]int{2, 4}), 4)
	forEachArray(7, []int{0, 1, 2, 3}, func(a []int) {
		expected := 0
		for i, h := range a {
			left, right := i, i
			for left > 0 && a[left-1] >= h {
				left--
			}
			for right+1 < len(a) && a[right+1] >= h {
				right++
			}
			if area := h * (right - left + 1); area > expected {
				expected = area
			}
		}
		eq(t, largestRectangleArea(a), expected)
	})
	eq(t, largestRectangleArea([]int{10000, 10000, 10000}), 30000)
	t.Logf("checks=%d", checks)
}
