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
		name  string
		n     int
		edges [][]int
		want  bool
	}{{"official_true", 2, [][]int{{1, 0}}, true}, {"official_false", 2, [][]int{{1, 0}, {0, 1}}, false}, {"disconnected_cycle", 4, [][]int{{1, 0}, {2, 3}, {3, 2}}, false}, {"self_cycle", 1, [][]int{{0, 0}}, false}, {"no_edges", 4, [][]int{}, true}, {"diamond", 4, [][]int{{1, 0}, {2, 0}, {3, 1}, {3, 2}}, true}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, canFinish(c.n, c.edges), c.want)
		})
	}
}
