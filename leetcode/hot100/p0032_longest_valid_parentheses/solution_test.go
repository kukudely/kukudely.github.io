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
	eachArray([]int{40, 41}, 12, func(a []int) {
		b := make([]byte, len(a))
		for i, x := range a {
			b[i] = byte(x)
		}
		s := string(b)
		want := 0
		for i := range b {
			balance := 0
			for j := i; j < len(b); j++ {
				if b[j] == '(' {
					balance++
				} else {
					balance--
				}
				if balance < 0 {
					break
				}
				if balance == 0 {
					want = max(want, j-i+1)
				}
			}
		}
		if got := longestValidParentheses(s); got != want {
			t.Fatalf("%q got %d want %d", s, got, want)
		}
	})
}
