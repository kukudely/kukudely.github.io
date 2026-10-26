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
func sortedArrays(maxLength int, alphabet []int) [][]int {
	answer := [][]int{}
	var walk func([]int, int)
	walk = func(a []int, start int) {
		answer = append(answer, append([]int{}, a...))
		if len(a) == maxLength {
			return
		}
		for i := start; i < len(alphabet); i++ {
			next := append(append([]int{}, a...), alphabet[i])
			walk(next, i)
		}
	}
	walk([]int{}, 0)
	return answer
}
func TestExamplesAndBoundaries(t *testing.T) {
	defer testutil.SkipTODO(t)
	eq(t, findMedianSortedArrays([]int{1, 3}, []int{2}), 2.0)
	eq(t, findMedianSortedArrays([]int{1, 2}, []int{3, 4}), 2.5)
	arrays := sortedArrays(4, []int{-2, -1, 0, 1, 2})
	for _, a := range arrays {
		for _, b := range arrays {
			if len(a)+len(b) == 0 {
				continue
			}
			joined := append(append([]int{}, a...), b...)
			sort.Ints(joined)
			n := len(joined)
			expected := float64(joined[n/2])
			if n%2 == 0 {
				expected = (float64(joined[n/2-1]) + float64(joined[n/2])) / 2
			}
			eq(t, findMedianSortedArrays(a, b), expected)
		}
	}
	eq(t, findMedianSortedArrays([]int{-1000000}, []int{1000000}), 0.0)
	t.Logf("checks=%d", checks)
}
