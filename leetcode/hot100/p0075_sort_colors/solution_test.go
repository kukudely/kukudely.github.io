package solution

import (
	"leetcodepractice/internal/testutil"
	"slices"
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
func sameInts(a, b []int) bool {
	return slices.Equal(a, b)
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	eachArray([]int{0, 1, 2}, 8, func(a []int) {
		got, want := slices.Clone(a), slices.Clone(a)
		slices.Sort(want)
		sortColors(got)
		if !sameInts(got, want) {
			t.Fatal(a, got, want)
		}
	})
}
