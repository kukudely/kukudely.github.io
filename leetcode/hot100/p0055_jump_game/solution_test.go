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
	eq(t, canJump([]int{2, 3, 1, 1, 4}), true)
	eq(t, canJump([]int{3, 2, 1, 0, 4}), false)
	forEachArray(7, []int{0, 1, 2, 3}, func(a []int) {
		expected := false
		if len(a) > 0 {
			visited := make([]bool, len(a))
			queue := []int{0}
			visited[0] = true
			for len(queue) > 0 {
				i := queue[0]
				queue = queue[1:]
				for next := i + 1; next < len(a) && next <= i+a[i]; next++ {
					if !visited[next] {
						visited[next] = true
						queue = append(queue, next)
					}
				}
			}
			expected = visited[len(a)-1]
		}
		eq(t, canJump(a), expected)
	})
	t.Logf("checks=%d", checks)
}
