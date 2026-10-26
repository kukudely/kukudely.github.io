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
func forEachArray(maxLength int, alphabet []int, visit func([]int)) {
	var walk func([]int)
	walk = func(a []int) {
		visit(a)
		if len(a) == maxLength {
			return
		}
		for _, value := range alphabet {
			next := append(append([]int{}, a...), value)
			walk(next)
		}
	}
	walk([]int{})
}
func TestExamplesAndBoundaries(t *testing.T) {
	defer testutil.SkipTODO(t)
	eq(t, jump([]int{2, 3, 1, 1, 4}), 2)
	eq(t, jump([]int{2, 3, 0, 1, 4}), 2)
	forEachArray(7, []int{0, 1, 2, 3}, func(a []int) {
		expected := 0
		if len(a) > 1 {
			distance := make([]int, len(a))
			for i := 1; i < len(a); i++ {
				distance[i] = -1
			}
			queue := []int{0}
			for len(queue) > 0 {
				i := queue[0]
				queue = queue[1:]
				for next := i + 1; next < len(a) && next <= i+a[i]; next++ {
					if distance[next] == -1 {
						distance[next] = distance[i] + 1
						queue = append(queue, next)
					}
				}
			}
			expected = distance[len(a)-1]
		}
		eq(t, jump(a), expected)
	})
	t.Logf("checks=%d", checks)
}
