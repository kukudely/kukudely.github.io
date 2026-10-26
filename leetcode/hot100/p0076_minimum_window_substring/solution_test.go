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
	equal(t, minWindow("ADOBECODEBANC", "ABC"), "BANC")
	equal(t, minWindow("aa", "aa"), "aa")
	equal(t, minWindow("a", "aa"), "")
	equal(t, minWindow("", "a"), "")
	rng := rand.New(rand.NewSource(76))
	for z := 0; z < 300; z++ {
		n, m := rng.Intn(12), 1+rng.Intn(5)
		s, p := make([]byte, n), make([]byte, m)
		for i := range s {
			s[i] = byte('A' + rng.Intn(3))
		}
		var need [128]int
		for i := range p {
			p[i] = byte('A' + rng.Intn(3))
			need[p[i]]++
		}
		best := ""
		for l := 0; l < n; l++ {
			for r := l + 1; r <= n; r++ {
				var have [128]int
				for _, c := range s[l:r] {
					have[c]++
				}
				ok := true
				for c := range need {
					if have[c] < need[c] {
						ok = false
						break
					}
				}
				if ok && (best == "" || r-l < len(best)) {
					best = string(s[l:r])
				}
			}
		}
		equal(t, minWindow(string(s), string(p)), best)
	}
}
