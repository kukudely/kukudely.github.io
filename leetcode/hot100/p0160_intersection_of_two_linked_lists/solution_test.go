package solution

import (
	"leetcodepractice/internal/testutil"
	"reflect"
	"testing"
)

func list(values ...int) *ListNode {
	dummy := &ListNode{}
	tail := dummy
	for _, x := range values {
		tail.Next = &ListNode{Val: x}
		tail = tail.Next
	}
	return dummy.Next
}
func values(t *testing.T, head *ListNode) []int {
	t.Helper()
	out := []int{}
	seen := map[*ListNode]bool{}
	for p := head; p != nil; p = p.Next {
		if seen[p] {
			t.Fatal("unexpected list cycle")
		}
		seen[p] = true
		out = append(out, p.Val)
	}
	return out
}
func equal(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got %#v, want %#v", got, want)
	}
}
func listNodes(p *ListNode) []*ListNode {
	out := []*ListNode{}
	for ; p != nil; p = p.Next {
		out = append(out, p)
	}
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	for m := 0; m <= 8; m++ {
		for n := 0; n <= 8; n++ {
			for shared := 0; shared <= 6; shared++ {
				av, bv, cv := make([]int, m), make([]int, n), make([]int, shared)
				a, b, c := list(av...), list(bv...), list(cv...)
				connect := func(p, tail *ListNode) *ListNode {
					if p == nil {
						return tail
					}
					q := p
					for q.Next != nil {
						q = q.Next
					}
					q.Next = tail
					return p
				}
				a, b = connect(a, c), connect(b, c)
				beforeA, beforeB := listNodes(a), listNodes(b)
				if got := getIntersectionNode(a, b); got != c {
					t.Fatal("incorrect intersection identity")
				}
				equal(t, listNodes(a), beforeA)
				equal(t, listNodes(b), beforeB)
			}
		}
	}
}
