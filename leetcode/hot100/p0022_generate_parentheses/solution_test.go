package solution

import (
	"fmt"
	"leetcodepractice/internal/testutil"
	"reflect"
	"sort"
	"testing"
)

func eq(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v; want=%v", got, want)
	}
}
func stringsKey(values []string) []string {
	result := append([]string{}, values...)
	sort.Strings(result)
	return result
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	t.Run("official", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		eq(t, stringsKey(generateParenthesis(3)), stringsKey([]string{"((()))", "(()())", "(())()", "()(())", "()()()"}))
	})
	counts := []int{1, 1, 2, 5, 14, 42, 132, 429, 1430}
	for n, want := range counts {
		t.Run(fmt.Sprintf("catalan_%d", n), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			got := stringsKey(generateParenthesis(n))
			eq(t, len(got), want)
			for i, s := range got {
				if i > 0 && s == got[i-1] {
					t.Fatal("duplicate")
				}
				eq(t, len(s), 2*n)
				balance := 0
				for _, ch := range s {
					if ch == '(' {
						balance++
					} else {
						balance--
					}
					if balance < 0 {
						t.Fatal("invalid prefix")
					}
				}
				eq(t, balance, 0)
			}
		})
	}
}
