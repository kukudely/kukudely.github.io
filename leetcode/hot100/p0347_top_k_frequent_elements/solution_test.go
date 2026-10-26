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
	normalize := func(a []int) []int {
		sort.Ints(a)
		return a
	}
	eq(t, normalize(topKFrequent([]int{1, 1, 1, 2, 2, 3}, 2)), []int{1, 2})
	eq(t, normalize(topKFrequent([]int{1}, 1)), []int{1})
	forEachArray(7, []int{-1, 0, 1}, func(a []int) {
		if len(a) == 0 {
			return
		}
		counts := map[int]int{}
		for _, x := range a {
			counts[x]++
		}
		type pair struct{ v, c int }
		ordered := []pair{}
		for v, c := range counts {
			ordered = append(ordered, pair{v, c})
		}
		sort.Slice(ordered, func(i, j int) bool {
			return ordered[i].c > ordered[j].c
		})
		for k := 1; k <= len(ordered); k++ {
			if k < len(ordered) && ordered[k-1].c == ordered[k].c {
				continue
			}
			expected := []int{}
			for i := 0; i < k; i++ {
				expected = append(expected, ordered[i].v)
			}
			eq(t, normalize(topKFrequent(a, k)), normalize(expected))
		}
	})
	t.Logf("checks=%d", checks)
}
