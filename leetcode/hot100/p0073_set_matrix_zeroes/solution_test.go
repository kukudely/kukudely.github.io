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
func randomArray(rng *rand.Rand, n, radius int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rng.Intn(2*radius+1) - radius
	}
	return a
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	a := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(a)
	equal(t, a, [][]int{{1, 0, 1}, {0, 0, 0}, {1, 0, 1}})
	rng := rand.New(rand.NewSource(73))
	for m := 1; m <= 8; m++ {
		for n := 1; n <= 8; n++ {
			for z := 0; z < 5; z++ {
				a := make([][]int, m)
				for i := range a {
					a[i] = randomArray(rng, n, 2)
				}
				want := grid(a)
				for i := range a {
					for j, x := range a[i] {
						if x == 0 {
							for r := 0; r < m; r++ {
								want[r][j] = 0
							}
							for c := 0; c < n; c++ {
								want[i][c] = 0
							}
						}
					}
				}
				setZeroes(a)
				equal(t, a, want)
			}
		}
	}
}
