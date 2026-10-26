package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(136))
	for unique := -20; unique <= 20; unique++ {
		a := []int{unique}
		for value := -20; value <= 20; value++ {
			if value != unique {
				a = append(a, value, value)
			}
		}
		rng.Shuffle(len(a), func(i, j int) {
			a[i], a[j] = a[j], a[i]
		})
		if got := singleNumber(a); got != unique {
			t.Fatal(unique, got)
		}
	}
}
