package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func equalInts(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
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
	equalInts(t, spiralOrder([][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}), []int{1, 2, 3, 6, 9, 8, 7, 4, 5})
	for m := 1; m <= 9; m++ {
		for n := 1; n <= 9; n++ {
			a := make([][]int, m)
			for i := range a {
				a[i] = make([]int, n)
				for j := range a[i] {
					a[i][j] = i*n + j
				}
			}
			cur := grid(a)
			want := []int{}
			for len(cur) > 0 && len(cur[0]) > 0 {
				want = append(want, cur[0]...)
				cur = cur[1:]
				if len(cur) == 0 {
					break
				}
				h, w := len(cur), len(cur[0])
				next := make([][]int, w)
				for i := 0; i < w; i++ {
					next[i] = make([]int, h)
					for j := 0; j < h; j++ {
						next[i][j] = cur[j][w-1-i]
					}
				}
				cur = next
			}
			equalInts(t, spiralOrder(a), want)
		}
	}
}
