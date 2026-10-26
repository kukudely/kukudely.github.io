package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"strconv"
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
	for _, c := range []struct{ in, want string }{{"3[a]2[bc]", "aaabcbc"}, {"3[a2[c]]", "accaccacc"}, {"2[abc]3[cd]ef", "abcabccdcdcdef"}, {"abc3[cd]xyz", "abccdcdcdxyz"}, {"10[a]", "aaaaaaaaaa"}, {"1[a]", "a"}, {"abc", "abc"}, {"", ""}, {"2[2[b]a]", "bbabba"}, {"12[z]", "zzzzzzzzzzzz"}} {
		eq(t, decodeString(c.in), c.want)
	}
	for depth := 0; depth <= 10; depth++ {
		encoded, expected := "ab", "ab"
		for d := 0; d < depth; d++ {
			encoded = "2[" + encoded + "]"
			expected = expected + expected
		}
		eq(t, decodeString(encoded), expected)
	}
	for a := 1; a <= 5; a++ {
		for b := 1; b <= 5; b++ {
			encoded := strconv.Itoa(a) + "[x" + strconv.Itoa(b) + "[yz]]q"
			unit := "x"
			for i := 0; i < b; i++ {
				unit += "yz"
			}
			expected := ""
			for i := 0; i < a; i++ {
				expected += unit
			}
			expected += "q"
			eq(t, decodeString(encoded), expected)
		}
	}
	t.Logf("checks=%d", checks)
}
