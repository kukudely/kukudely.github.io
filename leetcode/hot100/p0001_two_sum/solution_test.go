package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func clone(a []int) []int {
	return append([]int(nil), a...)
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		a      []int
		target int
	}{{[]int{2, 7, 11, 15}, 9}, {[]int{3, 2, 4}, 6}, {[]int{3, 3}, 6}, {[]int{-3, 4, 3, 90}, 0}, {[]int{0, 4, 3, 0}, 0}}
	for _, tc := range cases {
		got := twoSum(clone(tc.a), tc.target)
		if len(got) != 2 || got[0] == got[1] || got[0] < 0 || got[1] < 0 || got[0] >= len(tc.a) || got[1] >= len(tc.a) || tc.a[got[0]]+tc.a[got[1]] != tc.target {
			t.Fatalf("bad result %v", got)
		}
	}
}
