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
func TestExamplesAndBoundaries(t *testing.T) {
	defer testutil.SkipTODO(t)
	eq(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 0), 4)
	eq(t, search([]int{4, 5, 6, 7, 0, 1, 2}, 3), -1)
	eq(t, search(nil, 0), -1)
	for n := 1; n <= 30; n++ {
		for rotation := 0; rotation < n; rotation++ {
			a := make([]int, n)
			for i := range a {
				a[i] = ((i+rotation)%n)*2 - 10
			}
			for target := -12; target <= 2*n-8; target++ {
				expected := -1
				for i, x := range a {
					if x == target {
						expected = i
					}
				}
				eq(t, search(a, target), expected)
			}
		}
	}
	t.Logf("checks=%d", checks)
}
