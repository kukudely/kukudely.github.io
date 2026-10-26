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
	eq(t, partitionLabels("ababcbacadefegdehijhklij"), []int{9, 7, 8})
	eq(t, partitionLabels("eccbbbbdec"), []int{10})
	forEachArray(7, []int{0, 1, 2}, func(a []int) {
		bytes := make([]byte, len(a))
		for i, v := range a {
			bytes[i] = byte('a' + v)
		}
		s := string(bytes)
		expected := []int{}
		start := 0
		for end := 0; end < len(s); end++ {
			valid := true
			for i := start; i <= end; i++ {
				for j := end + 1; j < len(s); j++ {
					if s[i] == s[j] {
						valid = false
					}
				}
			}
			if valid {
				expected = append(expected, end-start+1)
				start = end + 1
			}
		}
		eq(t, partitionLabels(s), expected)
	})
	t.Logf("checks=%d", checks)
}
