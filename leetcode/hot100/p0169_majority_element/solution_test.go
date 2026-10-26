package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(169))
	for n := 1; n <= 50; n++ {
		for trial := 0; trial < 20; trial++ {
			a := make([]int, n)
			want := rng.Intn(101) - 50
			for i := range a {
				if i <= n/2 {
					a[i] = want
				} else {
					a[i] = want + 1 + rng.Intn(10)
				}
			}
			rng.Shuffle(n, func(i, j int) {
				a[i], a[j] = a[j], a[i]
			})
			if got := majorityElement(a); got != want {
				t.Fatal(a, got, want)
			}
		}
	}
}
