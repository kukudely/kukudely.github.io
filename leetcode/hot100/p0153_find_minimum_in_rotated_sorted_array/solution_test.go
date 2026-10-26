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
	eq(t, findMin([]int{3, 4, 5, 1, 2}), 1)
	eq(t, findMin([]int{4, 5, 6, 7, 0, 1, 2}), 0)
	eq(t, findMin([]int{11, 13, 15, 17}), 11)
	for n := 1; n <= 60; n++ {
		for r := 0; r < n; r++ {
			a := make([]int, n)
			for i := range a {
				a[i] = ((i+r)%n)*3 - 50
			}
			eq(t, findMin(a), -50)
		}
	}
	t.Logf("checks=%d", checks)
}
