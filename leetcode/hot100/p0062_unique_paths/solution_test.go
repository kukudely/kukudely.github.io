package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	for m := 1; m <= 10; m++ {
		for n := 1; n <= 10; n++ {
			var ref func(int, int) int
			ref = func(a, b int) int {
				if a == 1 || b == 1 {
					return 1
				}
				return ref(a-1, b) + ref(a, b-1)
			}
			if got, want := uniquePaths(m, n), ref(m, n); got != want {
				t.Fatalf("%d %d got %d want %d", m, n, got, want)
			}
		}
	}
}
