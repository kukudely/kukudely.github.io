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
		name string
		nums []int
		want [][]int
	}{{"official", []int{1, 2, 3}, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}}, {"single", []int{1}, [][]int{{1}}}, {"negative_unsorted", []int{2, -1}, [][]int{{2, -1}, {-1, 2}}}, {"empty_extension", []int{}, [][]int{{}}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			before := append([]int{}, c.nums...)
			eq(t, rowsKey(permute(c.nums), false), rowsKey(c.want, false))
			eq(t, c.nums, before)
		})
	}
}
