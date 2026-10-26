package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"sort"
	"strings"
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
	cases := []struct {
		name, digits string
		want         []string
	}{{"official", "23", []string{"ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"}}, {"single", "2", []string{"a", "b", "c"}}, {"four_letters", "7", []string{"p", "q", "r", "s"}}, {"empty_extension", "", []string{}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, stringsKey(letterCombinations(c.digits)), stringsKey(c.want))
		})
	}
	t.Run("many_choices", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		got := stringsKey(letterCombinations("79"))
		eq(t, len(got), 16)
		for _, s := range got {
			if len(s) != 2 || !strings.ContainsRune("pqrs", rune(s[0])) || !strings.ContainsRune("wxyz", rune(s[1])) {
				t.Fatal(s)
			}
		}
		for i := 1; i < len(got); i++ {
			if got[i] == got[i-1] {
				t.Fatal("duplicate")
			}
		}
	})
}
