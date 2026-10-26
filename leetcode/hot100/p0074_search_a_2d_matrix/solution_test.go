package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"testing"
)

var checks int

func eq(t *testing.T, got, want any) {
	t.Helper()
	checks++
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v (check %d)", got, want, checks)
	}
}
func TestExamplesAndBoundaries(t *testing.T) {
	defer testutil.SkipTODO(t)
	matrix := [][]int{{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}}
	eq(t, searchMatrix(matrix, 3), true)
	eq(t, searchMatrix(matrix, 13), false)
	eq(t, searchMatrix(nil, 3), false)
	eq(t, searchMatrix([][]int{{}}, 3), false)
	for rows := 1; rows <= 8; rows++ {
		for cols := 1; cols <= 8; cols++ {
			a := make([][]int, rows)
			for i := range a {
				a[i] = make([]int, cols)
				for j := range a[i] {
					a[i][j] = (i*cols+j)*2 - 20
				}
			}
			for target := -22; target < rows*cols*2-17; target++ {
				expected := false
				for _, row := range a {
					for _, x := range row {
						if x == target {
							expected = true
						}
					}
				}
				eq(t, searchMatrix(a, target), expected)
			}
		}
	}
	eq(t, searchMatrix([][]int{{1, 1, 2}, {3, 3, 4}}, 3), true)
	t.Logf("checks=%d", checks)
}
