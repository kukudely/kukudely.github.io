package solution

import (
	"fmt"
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
	t.Run("official", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		want := [][]string{{".Q..", "...Q", "Q...", "..Q."}, {"..Q.", "Q...", "...Q", ".Q.."}}
		eq(t, stringRowsKey(solveNQueens(4)), stringRowsKey(want))
	})
	counts := []int{0, 1, 0, 0, 2, 10, 4, 40, 92, 352}
	for n := 1; n <= 9; n++ {
		t.Run(fmt.Sprint(n), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			boards := solveNQueens(n)
			eq(t, len(boards), counts[n])
			keys := stringRowsKey(boards)
			for i := 1; i < len(keys); i++ {
				if keys[i] == keys[i-1] {
					t.Fatal("duplicate board")
				}
			}
			for _, board := range boards {
				eq(t, len(board), n)
				cols, d1, d2 := map[int]bool{}, map[int]bool{}, map[int]bool{}
				for r, line := range board {
					eq(t, len(line), n)
					count := 0
					for c, ch := range line {
						if ch == '.' {
							continue
						}
						if ch != 'Q' {
							t.Fatal("invalid char")
						}
						count++
						if cols[c] || d1[r-c] || d2[r+c] {
							t.Fatal("attacking queens")
						}
						cols[c], d1[r-c], d2[r+c] = true, true, true
					}
					eq(t, count, 1)
				}
			}
		})
	}
}
