package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	a, b := 1, 1
	for n := 0; n <= 45; n++ {
		if got := climbStairs(n); got != a {
			t.Fatalf("n %d got %d want %d", n, got, a)
		}
		a, b = b, a+b
	}
}
