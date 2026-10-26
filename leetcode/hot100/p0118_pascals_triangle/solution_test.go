package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	for n := 0; n <= 30; n++ {
		rows := generate(n)
		if len(rows) != n {
			t.Fatal(n)
		}
		for i, row := range rows {
			if len(row) != i+1 || row[0] != 1 || row[i] != 1 {
				t.Fatal(rows)
			}
			for j := 1; j < i; j++ {
				if row[j] != rows[i-1][j-1]+rows[i-1][j] {
					t.Fatal(rows)
				}
			}
		}
	}
}
