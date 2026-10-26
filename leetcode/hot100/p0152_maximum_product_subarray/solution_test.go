package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func eachArray(values []int, limit int, visit func([]int)) {
	var walk func([]int)
	walk = func(a []int) {
		visit(a)
		if len(a) == limit {
			return
		}
		for _, x := range values {
			walk(append(a, x))
		}
	}
	walk(nil)
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	eachArray([]int{-2, -1, 0, 1, 2}, 5, func(a []int) {
		if len(a) == 0 {
			return
		}
		want := a[0]
		for i := range a {
			product := 1
			for j := i; j < len(a); j++ {
				product *= a[j]
				want = max(want, product)
			}
		}
		if got := maxProduct(a); got != want {
			t.Fatalf("%v got %d want %d", a, got, want)
		}
	})
}
