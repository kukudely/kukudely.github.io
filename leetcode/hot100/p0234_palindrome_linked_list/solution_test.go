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
	for n := 0; n <= 10; n++ {
		for mask := 0; mask < (1 << n); mask++ {
			a := make([]int, n)
			for i := range a {
				a[i] = (mask >> i) & 1
			}
			head := list(a...)
			nodes := listNodes(head)
			want := true
			for l, r := 0, n-1; l < r; l, r = l+1, r-1 {
				if a[l] != a[r] {
					want = false
				}
			}
			equal(t, isPalindrome(head), want)
			equalInts(t, values(t, head), a)
			after := listNodes(head)
			if len(after) != len(nodes) {
				t.Fatal("structure changed")
			}
			for i := range nodes {
				if nodes[i] != after[i] {
					t.Fatal("list not restored")
				}
			}
		}
	}
}
