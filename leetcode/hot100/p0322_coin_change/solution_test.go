package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	for _, coins := range [][]int{{1, 2, 5}, {2}, {3, 7}, {4, 6}, {1}, {}} {
		for amount := 0; amount <= 50; amount++ {
			dp := make([]int, amount+1)
			for x := 1; x <= amount; x++ {
				dp[x] = 999
				for _, c := range coins {
					if c <= x {
						dp[x] = min(dp[x], dp[x-c]+1)
					}
				}
			}
			want := dp[amount]
			if want >= 999 {
				want = -1
			}
			if got := coinChange(coins, amount); got != want {
				t.Fatalf("%v %d got %d want %d", coins, amount, got, want)
			}
		}
	}
}
