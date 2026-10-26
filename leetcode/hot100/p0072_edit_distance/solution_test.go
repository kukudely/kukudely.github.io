package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

func eachArray(values []int, limit int, visit func([]int)) {
	var walk func([]int)
	walk = func(a []int) {
		visit(a)
		if len(a) == limit {
			return
		}
		for _, x := range values {
			walk(append(a, x))
		}
	}
	walk(nil)
}
func TestExamples(t *testing.T) {
	defer testutil.SkipTODO(t)
	inputs := []string{}
	eachArray([]int{97, 98}, 4, func(a []int) {
		b := make([]byte, len(a))
		for i, x := range a {
			b[i] = byte(x)
		}
		inputs = append(inputs, string(b))
	})
	var ref func(string, string) int
	ref = func(a, b string) int {
		if len(a) == 0 {
			return len(b)
		}
		if len(b) == 0 {
			return len(a)
		}
		if a[0] == b[0] {
			return ref(a[1:], b[1:])
		}
		return 1 + min(ref(a[1:], b), ref(a, b[1:]), ref(a[1:], b[1:]))
	}
	for _, a := range inputs {
		for _, b := range inputs {
			if got, want := minDistance(a, b), ref(a, b); got != want {
				t.Fatalf("%q %q got %d want %d", a, b, got, want)
			}
		}
	}
	if minDistance("horse", "ros") != 3 || minDistance("intention", "execution") != 5 {
		t.Fatal("examples")
	}
}
