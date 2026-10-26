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
	eq(t, maxProfit([]int{7, 1, 5, 3, 6, 4}), 5)
	eq(t, maxProfit([]int{7, 6, 4, 3, 1}), 0)
	forEachArray(7, []int{0, 1, 2, 3}, func(a []int) {
		expected := 0
		for buy := range a {
			for sell := buy + 1; sell < len(a); sell++ {
				if a[sell]-a[buy] > expected {
					expected = a[sell] - a[buy]
				}
			}
		}
		eq(t, maxProfit(a), expected)
	})
	t.Logf("checks=%d", checks)
}
