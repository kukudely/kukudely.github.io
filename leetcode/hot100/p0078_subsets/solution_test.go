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
	}{{"example", []int{1, 2}, [][]int{{}, {1}, {2}, {1, 2}}}, {"single", []int{0}, [][]int{{}, {0}}}, {"negative", []int{-1, 2}, [][]int{{}, {-1}, {2}, {-1, 2}}}, {"empty_extension", []int{}, [][]int{{}}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, rowsKey(subsets(c.nums), true), rowsKey(c.want, true))
		})
	}
	t.Run("ten_elements_count_unique", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		rows := subsets([]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})
		keys := rowsKey(rows, true)
		eq(t, len(keys), 1024)
		for i := 1; i < len(keys); i++ {
			if keys[i] == keys[i-1] {
				t.Fatal("duplicate subset")
			}
		}
	})
}
