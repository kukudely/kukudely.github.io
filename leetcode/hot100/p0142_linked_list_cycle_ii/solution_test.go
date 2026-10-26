package solution

import (
	"leetcodepractice/internal/testutil"
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
func listNodes(p *ListNode) []*ListNode {
	out := []*ListNode{}
	for ; p != nil; p = p.Next {
		out = append(out, p)
	}
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	for n := 0; n <= 20; n++ {
		for pos := -1; pos < n; pos++ {
			head := list(make([]int, n)...)
			nodes := listNodes(head)
			var want *ListNode
			if pos >= 0 {
				want = nodes[pos]
				nodes[n-1].Next = want
			}
			if got := detectCycle(head); got != want {
				t.Fatal("incorrect entry identity")
			}
			if n > 0 && nodes[n-1].Next != want {
				t.Fatal("input changed")
			}
		}
	}
}
