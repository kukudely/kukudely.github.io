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
func equalInts(t *testing.T, got, want []int) {
	t.Helper()
	if len(got) != len(want) {
		t.Fatalf("got %v, want %v", got, want)
	}
	for i := range got {
		if got[i] != want[i] {
			t.Fatalf("got %v, want %v", got, want)
		}
	}
}
func clone(a []int) []int {
	return append([]int(nil), a...)
}
func listNodes(p *ListNode) []*ListNode {
	out := []*ListNode{}
	for ; p != nil; p = p.Next {
		out = append(out, p)
	}
	return out
}
func assertNodes(t *testing.T, p *ListNode, nodes []*ListNode) {
	t.Helper()
	got := values(t, p)
	if len(got) != len(nodes) {
		t.Fatal("node count changed")
	}
	wanted := map[*ListNode]bool{}
	for _, n := range nodes {
		wanted[n] = true
	}
	for ; p != nil; p = p.Next {
		if !wanted[p] {
			t.Fatal("node identity changed")
		}
		delete(wanted, p)
	}
	if len(wanted) != 0 {
		t.Fatal("nodes lost")
	}
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equalInts(t, values(t, reverseKGroup(list(1, 2, 3, 4, 5), 2)), []int{2, 1, 4, 3, 5})
	for n := 1; n <= 30; n++ {
		for k := 1; k <= n; k++ {
			a := make([]int, n)
			for i := range a {
				a[i] = i
			}
			head := list(a...)
			nodes := listNodes(head)
			want := clone(a)
			for start := 0; start+k <= n; start += k {
				for l, r := start, start+k-1; l < r; l, r = l+1, r-1 {
					want[l], want[r] = want[r], want[l]
				}
			}
			got := reverseKGroup(head, k)
			equalInts(t, values(t, got), want)
			assertNodes(t, got, nodes)
			for i, p := range nodes {
				if p.Val != a[i] {
					t.Fatal("node values changed")
				}
			}
		}
	}
}
