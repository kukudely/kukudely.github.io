package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"reflect"
	"sort"
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
func randomArray(rng *rand.Rand, n, radius int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rng.Intn(2*radius+1) - radius
	}
	return a
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equal(t, firstMissingPositive([]int{3, 4, -1, 1}), 2)
	equal(t, firstMissingPositive([]int{1, 1}), 2)
	equal(t, firstMissingPositive([]int{1, 2, 3}), 4)
	rng := rand.New(rand.NewSource(41))
	for z := 0; z < 500; z++ {
		a := randomArray(rng, 1+rng.Intn(25), 15)
		sorted := clone(a)
		sort.Ints(sorted)
		want := 1
		for _, x := range sorted {
			if x == want {
				want++
			}
		}
		equal(t, firstMissingPositive(a), want)
	}
}
