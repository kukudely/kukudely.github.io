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
func stringRowsKey(values [][]string) []string {
	result := []string{}
	for _, row := range values {
		result = append(result, strings.Join(row, "|"))
	}
	sort.Strings(result)
	return result
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		name, s string
		want    [][]string
	}{{"official", "aab", [][]string{{"a", "a", "b"}, {"aa", "b"}}}, {"single", "a", [][]string{{"a"}}}, {"all_same", "aaa", [][]string{{"a", "a", "a"}, {"a", "aa"}, {"aa", "a"}, {"aaa"}}}, {"whole_palindrome", "aba", [][]string{{"a", "b", "a"}, {"aba"}}}, {"empty_extension", "", [][]string{{}}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, stringRowsKey(partition(c.s)), stringRowsKey(c.want))
		})
	}
}
