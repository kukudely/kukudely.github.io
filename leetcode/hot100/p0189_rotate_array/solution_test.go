package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func equalInts(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
func clone(a []int) []int {
	return append([]int(nil), a...)
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	a := []int{1, 2, 3, 4, 5, 6, 7}
	rotate(a, 3)
	equalInts(t, a, []int{5, 6, 7, 1, 2, 3, 4})
	for n := 1; n <= 25; n++ {
		for k := 0; k <= 2*n+2; k++ {
			a := make([]int, n)
			for i := range a {
				a[i] = i
			}
			want := clone(a)
			for step := 0; step < k; step++ {
				last := want[n-1]
				copy(want[1:], want[:n-1])
				want[0] = last
			}
			rotate(a, k)
			equalInts(t, a, want)
		}
	}
}
