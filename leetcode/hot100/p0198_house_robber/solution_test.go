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
	eachArray([]int{0, 1, 3}, 7, func(a []int) {
		want := 0
		for mask := 0; mask < 1<<len(a); mask++ {
			if mask&(mask<<1) != 0 {
				continue
			}
			sum := 0
			for i, x := range a {
				if mask>>i&1 == 1 {
					sum += x
				}
			}
			want = max(want, sum)
		}
		if got := rob(a); got != want {
			t.Fatalf("%v got %d want %d", a, got, want)
		}
	})
}
