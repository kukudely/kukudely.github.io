package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	ref := make([]int, 10001)
	for n := 1; n < len(ref); n++ {
		ref[n] = n
		for x := 1; x*x <= n; x++ {
			ref[n] = min(ref[n], ref[n-x*x]+1)
		}
	}
	for n := 0; n <= 100; n++ {
		if got := numSquares(n); got != ref[n] {
			t.Fatalf("%d got %d want %d", n, got, ref[n])
		}
	}
	for _, n := range []int{997, 9999, 10000} {
		if numSquares(n) != ref[n] {
			t.Fatal(n)
		}
	}
}
