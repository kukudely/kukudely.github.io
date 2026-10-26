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
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equal(t, searchMatrix([][]int{{1, 4, 7}, {2, 5, 8}, {3, 6, 9}}, 6), true)
	rng := rand.New(rand.NewSource(240))
	for m := 1; m <= 8; m++ {
		for n := 1; n <= 8; n++ {
			a := make([][]int, m)
			seen := map[int]bool{}
			for i := 0; i < m; i++ {
				a[i] = make([]int, n)
				for j := 0; j < n; j++ {
					base := -10
					if i > 0 {
						base = max(base, a[i-1][j])
					}
					if j > 0 {
						base = max(base, a[i][j-1])
					}
					a[i][j] = base + rng.Intn(3)
					seen[a[i][j]] = true
				}
			}
			for target := -11; target <= 30; target++ {
				equal(t, searchMatrix(a, target), seen[target])
			}
		}
	}
}
