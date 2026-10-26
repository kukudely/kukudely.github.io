package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"reflect"
	"sort"
	"testing"
)

func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
func clone(a []int) []int {
	return append([]int(nil), a...)
}
func grid(a [][]int) [][]int {
	out := make([][]int, len(a))
	for i := range a {
		out[i] = clone(a[i])
	}
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct{ a, w [][]int }{{[][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}}, [][]int{{1, 6}, {8, 10}, {15, 18}}}, {[][]int{{1, 4}, {4, 5}}, [][]int{{1, 5}}}, {[][]int{{1, 10}, {2, 3}, {4, 5}}, [][]int{{1, 10}}}, {[][]int{{2, 2}, {1, 2}, {5, 7}, {3, 5}}, [][]int{{1, 2}, {3, 7}}}, {[][]int{{1, 1}, {1, 1}}, [][]int{{1, 1}}}}
	for _, tc := range cases {
		equal(t, merge(grid(tc.a)), tc.w)
	}
	rng := rand.New(rand.NewSource(56))
	for z := 0; z < 200; z++ {
		n := 1 + rng.Intn(12)
		a := make([][]int, n)
		parent := make([]int, n)
		for i := range parent {
			parent[i] = i
		}
		var find func(int) int
		find = func(i int) int {
			for parent[i] != i {
				i = parent[i]
			}
			return i
		}
		for i := range a {
			x, y := rng.Intn(20), rng.Intn(20)
			a[i] = []int{min(x, y), max(x, y)}
		}
		for i := range a {
			for j := 0; j < i; j++ {
				if max(a[i][0], a[j][0]) <= min(a[i][1], a[j][1]) {
					parent[find(i)] = find(j)
				}
			}
		}
		groups := map[int][]int{}
		for i, p := range a {
			root := find(i)
			if g, ok := groups[root]; ok {
				g[0] = min(g[0], p[0])
				g[1] = max(g[1], p[1])
			} else {
				groups[root] = clone(p)
			}
		}
		want := [][]int{}
		for _, p := range groups {
			want = append(want, p)
		}
		sort.Slice(want, func(i, j int) bool {
			return want[i][0] < want[j][0]
		})
		equal(t, merge(grid(a)), want)
	}
}
