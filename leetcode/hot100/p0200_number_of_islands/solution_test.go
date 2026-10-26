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
func bytesGrid(rows ...string) [][]byte {
	grid := make([][]byte, len(rows))
	for i, row := range rows {
		grid[i] = []byte(row)
	}
	return grid
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		name string
		rows []string
		want int
	}{{"official", []string{"11000", "11000", "00100", "00011"}, 3}, {"diagonal", []string{"10", "01"}, 2}, {"water", []string{"000", "000"}, 0}, {"land", []string{"111", "111"}, 1}, {"single", []string{"1"}, 1}, {"empty", []string{}, 0}, {"empty_row", []string{""}, 0}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, numIslands(bytesGrid(c.rows...)), c.want)
		})
	}
}
