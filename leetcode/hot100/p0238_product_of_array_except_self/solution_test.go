package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
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
func randomArray(rng *rand.Rand, n, radius int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rng.Intn(2*radius+1) - radius
	}
	return a
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equalInts(t, productExceptSelf([]int{1, 2, 3, 4}), []int{24, 12, 8, 6})
	equalInts(t, productExceptSelf([]int{-1, 1, 0, -3, 3}), []int{0, 0, 9, 0, 0})
	equalInts(t, productExceptSelf([]int{0, 0, 2}), []int{0, 0, 0})
	rng := rand.New(rand.NewSource(238))
	for z := 0; z < 200; z++ {
		a := randomArray(rng, 2+rng.Intn(10), 2)
		want := make([]int, len(a))
		for i := range want {
			want[i] = 1
			for j, x := range a {
				if i != j {
					want[i] *= x
				}
			}
		}
		equalInts(t, productExceptSelf(a), want)
	}
}
