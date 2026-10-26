package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(287))
	for n := 1; n <= 50; n++ {
		for duplicate := 1; duplicate <= n; duplicate++ {
			a := []int{duplicate}
			for i := 1; i <= n; i++ {
				a = append(a, i)
			}
			rng.Shuffle(len(a), func(i, j int) {
				a[i], a[j] = a[j], a[i]
			})
			if got := findDuplicate(a); got != duplicate {
				t.Fatal(a, got, duplicate)
			}
		}
	}
	for trial := 0; trial < 300; trial++ {
		n := 1 + rng.Intn(50)
		duplicate := 1 + rng.Intn(n)
		times := 2 + rng.Intn(n)
		a := make([]int, times)
		for i := range a {
			a[i] = duplicate
		}
		for x := 1; len(a) < n+1; x++ {
			if x != duplicate {
				a = append(a, x)
			}
		}
		rng.Shuffle(len(a), func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
		if got := findDuplicate(a); got != duplicate {
			t.Fatal(a, got, duplicate)
		}
	}
}
