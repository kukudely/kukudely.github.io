package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"sort"
	"strings"
	"testing"
)

func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
func canonical(groups [][]string) []string {
	out := []string{}
	for _, g := range groups {
		cp := append([]string(nil), g...)
		sort.Strings(cp)
		out = append(out, strings.Join(cp, "|"))
	}
	sort.Strings(out)
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equal(t, canonical(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})), canonical([][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}))
	equal(t, canonical(groupAnagrams([]string{"", "", "a", "aa"})), canonical([][]string{{"", ""}, {"a"}, {"aa"}}))
	equal(t, canonical(groupAnagrams(nil)), []string{})
}
