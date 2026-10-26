package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
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
func listNodes(p *ListNode) []*ListNode {
	out := []*ListNode{}
	for ; p != nil; p = p.Next {
		out = append(out, p)
	}
	return out
}
func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	equalInts(t, values(t, addTwoNumbers(list(2, 4, 3), list(5, 6, 4))), []int{7, 0, 8})
	equalInts(t, values(t, addTwoNumbers(list(9, 9, 9), list(1))), []int{0, 0, 0, 1})
	toList := func(x int) *ListNode {
		if x == 0 {
			return list(0)
		}
		a := []int{}
		for x > 0 {
			a = append(a, x%10)
			x /= 10
		}
		return list(a...)
	}
	rng := rand.New(rand.NewSource(2))
	for z := 0; z < 300; z++ {
		a, b := rng.Intn(100000000), rng.Intn(100000000)
		la, lb := toList(a), toList(b)
		va, vb := values(t, la), values(t, lb)
		got := addTwoNumbers(la, lb)
		equalInts(t, values(t, got), values(t, toList(a+b)))
		equalInts(t, values(t, la), va)
		equalInts(t, values(t, lb), vb)
		original := map[*ListNode]bool{}
		for _, p := range append(listNodes(la), listNodes(lb)...) {
			original[p] = true
		}
		for p := got; p != nil; p = p.Next {
			if original[p] {
				t.Fatal("result aliases input")
			}
		}
	}
}
