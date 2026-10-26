package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"sort"
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
	f := Constructor()
	f.AddNum(1)
	f.AddNum(2)
	eq(t, f.FindMedian(), 1.5)
	f.AddNum(3)
	eq(t, f.FindMedian(), 2.0)
	for _, sequence := range [][]int{{1}, {-1, -2, -3, -4, -5}, {5, 4, 3, 2, 1}, {0, 0, 0, 0}, {-100000, 100000, 100000, -100000}} {
		finder := Constructor()
		reference := []int{}
		for _, x := range sequence {
			finder.AddNum(x)
			reference = append(reference, x)
			sort.Ints(reference)
			n := len(reference)
			want := float64(reference[n/2])
			if n%2 == 0 {
				want = (float64(reference[n/2-1]) + float64(reference[n/2])) / 2
			}
			eq(t, finder.FindMedian(), want)
		}
	}
	finder := Constructor()
	reference := []int{}
	seed := uint64(1729)
	for i := 0; i < 5000; i++ {
		seed = seed*6364136223846793005 + 1
		x := int(seed%200001) - 100000
		finder.AddNum(x)
		reference = append(reference, x)
		sort.Ints(reference)
		n := len(reference)
		want := float64(reference[n/2])
		if n%2 == 0 {
			want = (float64(reference[n/2-1]) + float64(reference[n/2])) / 2
		}
		eq(t, finder.FindMedian(), want)
	}
	t.Logf("checks=%d", checks)
}
