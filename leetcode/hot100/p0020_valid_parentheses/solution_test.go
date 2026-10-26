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
	for _, c := range []struct {
		s    string
		want bool
	}{{"()", true}, {"()[]{}", true}, {"(]", false}, {"([])", true}, {"([)]", false}, {"", true}, {"(((", false}, {"})", false}} {
		eq(t, isValid(c.s), c.want)
	}
	var oracle func(string) bool
	oracle = func(s string) bool {
		if len(s) == 0 {
			return true
		}
		close := byte(0)
		switch s[0] {
		case '(':
			close = ')'
		case '[':
			close = ']'
		case '{':
			close = '}'
		default:
			return false
		}
		for i := 1; i < len(s); i += 2 {
			if s[i] == close && oracle(s[1:i]) && oracle(s[i+1:]) {
				return true
			}
		}
		return false
	}
	var visit func(string, int)
	visit = func(s string, remaining int) {
		eq(t, isValid(s), oracle(s))
		if remaining == 0 {
			return
		}
		for _, ch := range "()[]{}" {
			visit(s+string(ch), remaining-1)
		}
	}
	visit("", 6)
	t.Logf("checks=%d", checks)
}
