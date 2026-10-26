package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"sort"
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
func randomArray(rng *rand.Rand, n, radius int) []int {
	a := make([]int, n)
	for i := range a {
		a[i] = rng.Intn(2*radius+1) - radius
	}
	return a
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
	rng := rand.New(rand.NewSource(21))
	for z := 0; z < 250; z++ {
		a, b := randomArray(rng, rng.Intn(20), 8), randomArray(rng, rng.Intn(20), 8)
		sort.Ints(a)
		sort.Ints(b)
		la, lb := list(a...), list(b...)
		nodes := append(listNodes(la), listNodes(lb)...)
		want := append(clone(a), b...)
		sort.Ints(want)
		got := mergeTwoLists(la, lb)
		equalInts(t, values(t, got), want)
		assertNodes(t, got, nodes)
	}
}
