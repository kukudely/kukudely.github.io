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
	eq(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 8), []int{3, 4})
	eq(t, searchRange([]int{5, 7, 7, 8, 8, 10}, 6), []int{-1, -1})
	for _, a := range sortedArrays(7, []int{-1, 0, 1}) {
		for target := -2; target <= 2; target++ {
			expected := []int{-1, -1}
			for i, x := range a {
				if x == target {
					if expected[0] < 0 {
						expected[0] = i
					}
					expected[1] = i
				}
			}
			eq(t, searchRange(a, target), expected)
		}
	}
	eq(t, searchRange([]int{1000000000, 1000000000}, 1000000000), []int{0, 1})
	t.Logf("checks=%d", checks)
}
