package solution

import (
	"leetcodepractice/internal/testutil"
	"strings"
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
func palindrome(s string) bool {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		if s[i] != s[j] {
			return false
		}
	}
	return true
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	eachArray([]int{97, 98, 99}, 6, func(a []int) {
		b := make([]byte, len(a))
		for i, x := range a {
			b[i] = byte(x)
		}
		s := string(b)
		want := 0
		for i := 0; i < len(s); i++ {
			for j := i + 1; j <= len(s); j++ {
				if palindrome(s[i:j]) {
					want = max(want, j-i)
				}
			}
		}
		got := longestPalindrome(s)
		if len(got) != want || !palindrome(got) || !strings.Contains(s, got) {
			t.Fatalf("%q got %q want length %d", s, got, want)
		}
	})
}
