package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"reflect"
	"testing"
)

func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
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
	equal(t, maxSubArray([]int{-2, 1, -3, 4, -1, 2, 1, -5, 4}), 6)
	equal(t, maxSubArray([]int{-5, -2, -8}), -2)
	rng := rand.New(rand.NewSource(53))
	for z := 0; z < 300; z++ {
		a := randomArray(rng, 1+rng.Intn(20), 10)
		want := a[0]
		for l := range a {
			for r := l + 1; r <= len(a); r++ {
				sum := 0
				for _, x := range a[l:r] {
					sum += x
				}
				want = max(want, sum)
			}
		}
		equal(t, maxSubArray(a), want)
	}
}
