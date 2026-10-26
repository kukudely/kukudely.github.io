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
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	dictionaries := [][]string{{"a", "ab", "bc"}, {"a", "b", "c"}, {"abc"}, {"aa", "ab", "ba"}, {"leet", "code"}}
	for _, words := range dictionaries {
		eachArray([]int{97, 98, 99}, 5, func(a []int) {
			b := make([]byte, len(a))
			for i, x := range a {
				b[i] = byte(x)
			}
			s := string(b)
			var ref func(string) bool
			ref = func(s string) bool {
				if len(s) == 0 {
					return true
				}
				for _, w := range words {
					if strings.HasPrefix(s, w) && ref(s[len(w):]) {
						return true
					}
				}
				return false
			}
			if got, want := wordBreak(s, words), ref(s); got != want {
				t.Fatalf("%q %v got %v", s, words, got)
			}
		})
	}
	if !wordBreak("leetcode", []string{"leet", "code"}) || wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"}) {
		t.Fatal("examples")
	}
}
