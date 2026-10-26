package solution

import (
	"fmt"
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
func nextReference(a []int) []int {
	if len(a) == 0 {
		return nil
	}
	all := [][]int{}
	seen := map[string]bool{}
	var walk func(int)
	work := slices.Clone(a)
	walk = func(i int) {
		if i == len(work) {
			key := fmt.Sprint(work)
			if !seen[key] {
				seen[key] = true
				all = append(all, slices.Clone(work))
			}
			return
		}
		for j := i; j < len(work); j++ {
			work[i], work[j] = work[j], work[i]
			walk(i + 1)
			work[i], work[j] = work[j], work[i]
		}
	}
	walk(0)
	slices.SortFunc(all, func(a, b []int) int {
		return slices.Compare(a, b)
	})
	for i, row := range all {
		if slices.Equal(row, a) {
			return all[(i+1)%len(all)]
		}
	}
	panic("missing")
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	eachArray([]int{0, 1, 2}, 5, func(a []int) {
		if len(a) == 0 {
			return
		}
		want := nextReference(a)
		got := slices.Clone(a)
		nextPermutation(got)
		if !sameInts(got, want) {
			t.Fatal(a, got, want)
		}
	})
}
