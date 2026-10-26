package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(300))
	for trial := 0; trial < 500; trial++ {
		a := make([]int, rng.Intn(11))
		for i := range a {
			a[i] = rng.Intn(9) - 4
		}
		want := 0
		for mask := 0; mask < 1<<len(a); mask++ {
			last, count, valid := 0, 0, true
			for i, x := range a {
				if mask>>i&1 == 1 {
					if count > 0 && x <= last {
						valid = false
					}
					last = x
					count++
				}
			}
			if valid {
				want = max(want, count)
			}
		}
		if got := lengthOfLIS(a); got != want {
			t.Fatalf("%v got %d want %d", a, got, want)
		}
	}
}
