package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"testing"
)

func eq(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v; want=%v", got, want)
	}
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		name string
		grid [][]int
		want int
	}{{"official", [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}, 4}, {"unreachable", [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}}, -1}, {"no_fresh", [][]int{{0, 2}}, 0}, {"no_source", [][]int{{1}}, -1}, {"multiple_sources", [][]int{{2, 1, 1, 1, 2}}, 2}, {"single_minute", [][]int{{2, 1}}, 1}, {"empty", [][]int{}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, orangesRotting(c.grid), c.want)
		})
	}
}
