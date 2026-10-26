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
func rowsKey(rows [][]int, sortInner bool) []string {
	result := []string{}
	for _, row := range rows {
		values := append([]int{}, row...)
		if sortInner {
			sort.Ints(values)
		}
		result = append(result, fmt.Sprint(values))
	}
	sort.Strings(result)
	return result
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{{"official", []int{2, 3, 6, 7}, 7, [][]int{{2, 2, 3}, {7}}}, {"many", []int{2, 3, 5}, 8, [][]int{{2, 2, 2, 2}, {2, 3, 3}, {3, 5}}}, {"none", []int{2}, 1, [][]int{}}, {"one_repeated", []int{2}, 6, [][]int{{2, 2, 2}}}, {"unsorted", []int{7, 3, 2}, 7, [][]int{{7}, {2, 2, 3}}}, {"zero_extension", []int{2}, 0, [][]int{{}}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			before := append([]int{}, c.candidates...)
			eq(t, rowsKey(combinationSum(c.candidates, c.target), true), rowsKey(c.want, true))
			eq(t, c.candidates, before)
		})
	}
}
