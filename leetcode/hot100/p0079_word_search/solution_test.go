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
		word string
		want bool
	}{{"official_true", []string{"ABCE", "SFCS", "ADEE"}, "ABCCED", true}, {"official_see", []string{"ABCE", "SFCS", "ADEE"}, "SEE", true}, {"no_reuse", []string{"ABCE", "SFCS", "ADEE"}, "ABCB", false}, {"single", []string{"A"}, "A", true}, {"too_long", []string{"A"}, "AA", false}, {"frequency_prune", []string{"AB", "CD"}, "AAA", false}, {"reverse_search", []string{"AAAA", "AAAB"}, "AAAB", true}, {"diagonal_forbidden", []string{"AX", "XB"}, "AB", false}, {"empty_word_extension", []string{"A"}, "", true}, {"empty_grid_extension", []string{}, "A", false}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			board := bytesGrid(c.rows...)
			eq(t, exist(board, c.word), c.want)
			eq(t, board, bytesGrid(c.rows...))
		})
	}
}
