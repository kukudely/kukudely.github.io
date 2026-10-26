package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"sort"
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
func randomArray(rng *rand.Rand, n, radius int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rng.Intn(2*radius+1) - radius
	}
	return a
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equalInts(t, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3), []int{3, 3, 5, 5, 6, 7})
	rng := rand.New(rand.NewSource(239))
	for n := 1; n <= 30; n++ {
		a := randomArray(rng, n, 5)
		for k := 1; k <= n; k++ {
			want := []int{}
			for l := 0; l+k <= n; l++ {
				cp := clone(a[l : l+k])
				sort.Ints(cp)
				want = append(want, cp[k-1])
			}
			equalInts(t, maxSlidingWindow(a, k), want)
		}
	}
}
