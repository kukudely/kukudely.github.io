package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	rng := rand.New(rand.NewSource(64))
	for trial := 0; trial < 300; trial++ {
		m, n := 1+rng.Intn(5), 1+rng.Intn(5)
		grid := make([][]int, m)
		for i := range grid {
			grid[i] = make([]int, n)
			for j := range grid[i] {
				grid[i][j] = rng.Intn(10)
			}
		}
		var ref func(int, int) int
		ref = func(i, j int) int {
			if i == m-1 && j == n-1 {
				return grid[i][j]
			}
			best := int(1e9)
			if i+1 < m {
				best = min(best, ref(i+1, j))
			}
			if j+1 < n {
				best = min(best, ref(i, j+1))
			}
			return grid[i][j] + best
		}
		want := ref(0, 0)
		if got := minPathSum(grid); got != want {
			t.Fatalf("got %d want %d", got, want)
		}
	}
}
