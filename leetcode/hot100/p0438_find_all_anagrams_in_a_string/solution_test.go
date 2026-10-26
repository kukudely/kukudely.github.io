package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func equalInts(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equalInts(t, findAnagrams("cbaebabacd", "abc"), []int{0, 6})
	equalInts(t, findAnagrams("abab", "ab"), []int{0, 1, 2})
	equalInts(t, findAnagrams("a", "aa"), nil)
	rng := rand.New(rand.NewSource(438))
	for z := 0; z < 300; z++ {
		n, m := rng.Intn(16), 1+rng.Intn(8)
		s, p := make([]byte, n), make([]byte, m)
		for i := range s {
			s[i] = byte('a' + rng.Intn(3))
		}
		var target [26]int
		for i := range p {
			p[i] = byte('a' + rng.Intn(3))
			target[p[i]-'a']++
		}
		want := []int{}
		for l := 0; l+m <= n; l++ {
			var count [26]int
			for _, c := range s[l : l+m] {
				count[c-'a']++
			}
			if count == target {
				want = append(want, l)
			}
		}
		equalInts(t, findAnagrams(string(s), string(p)), want)
	}
}
