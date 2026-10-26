package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func eachArray(values []int, limit int, visit func([]int)) {
	var walk func([]int)
	walk = func(a []int) {
		visit(a)
		if len(a) == limit {
			return
		}
		for _, x := range values {
			walk(append(a, x))
		}
	}
	walk(nil)
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	inputs := []string{}
	eachArray([]int{97, 98}, 4, func(a []int) {
		b := make([]byte, len(a))
		for i, x := range a {
			b[i] = byte(x)
		}
		inputs = append(inputs, string(b))
	})
	for _, a := range inputs {
		for _, b := range inputs {
			want := 0
			for mask := 0; mask < 1<<len(a); mask++ {
				sub := []byte{}
				for i := range a {
					if mask>>i&1 == 1 {
						sub = append(sub, a[i])
					}
				}
				p := 0
				for i := range b {
					if p < len(sub) && sub[p] == b[i] {
						p++
					}
				}
				if p == len(sub) {
					want = max(want, p)
				}
			}
			if got := longestCommonSubsequence(a, b); got != want {
				t.Fatalf("%q %q got %d want %d", a, b, got, want)
			}
		}
	}
}
