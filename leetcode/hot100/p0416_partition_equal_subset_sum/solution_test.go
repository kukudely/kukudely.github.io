package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(416))
	for trial := 0; trial < 500; trial++ {
		a := make([]int, 1+rng.Intn(10))
		total := 0
		for i := range a {
			a[i] = 1 + rng.Intn(10)
			total += a[i]
		}
		want := false
		for mask := 0; mask < 1<<len(a); mask++ {
			sum := 0
			for i, x := range a {
				if mask>>i&1 == 1 {
					sum += x
				}
			}
			if 2*sum == total {
				want = true
				break
			}
		}
		if got := canPartition(a); got != want {
			t.Fatalf("%v got %v want %v", a, got, want)
		}
	}
}
