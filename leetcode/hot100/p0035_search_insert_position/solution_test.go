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
	eq(t, searchInsert([]int{1, 3, 5, 6}, 5), 2)
	eq(t, searchInsert([]int{1, 3, 5, 6}, 2), 1)
	eq(t, searchInsert([]int{1, 3, 5, 6}, 7), 4)
	for mask := 0; mask < 256; mask++ {
		a := []int{}
		for bit := 0; bit < 8; bit++ {
			if mask&(1<<bit) != 0 {
				a = append(a, bit*2-7)
			}
		}
		for target := -9; target <= 10; target++ {
			expected := 0
			for _, x := range a {
				if x < target {
					expected++
				}
			}
			eq(t, searchInsert(a, target), expected)
		}
	}
	t.Logf("checks=%d", checks)
}
