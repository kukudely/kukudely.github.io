package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"testing"
)

func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
func clone(a []int) []int {
	return append([]int(nil), a...)
}
func grid(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = clone(a[i])
	}
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	a := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(a)
	equal(t, a, [][]int{{7, 4, 1}, {8, 5, 2}, {9, 6, 3}})
	for n := 1; n <= 15; n++ {
		a := make([][]int, n)
		want := make([][]int, n)
		for i := range a {
			a[i] = make([]int, n)
			want[i] = make([]int, n)
			for j := range a[i] {
				a[i][j] = i*n + j
			}
		}
		original := grid(a)
		for i := range a {
			for j := range a[i] {
				want[j][n-1-i] = a[i][j]
			}
		}
		rotate(a)
		equal(t, a, want)
		rotate(a)
		rotate(a)
		rotate(a)
		equal(t, a, original)
	}
}
