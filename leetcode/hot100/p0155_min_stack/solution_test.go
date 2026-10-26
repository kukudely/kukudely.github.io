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
	s := Constructor()
	s.Push(-2)
	s.Push(0)
	s.Push(-3)
	eq(t, s.GetMin(), -3)
	s.Pop()
	eq(t, s.Top(), 0)
	eq(t, s.GetMin(), -2)
	s.Pop()
	s.Pop()
	reference := []int{}
	seed := uint64(1729)
	for step := 0; step < 16000; step++ {
		seed = seed*6364136223846793005 + 1
		if len(reference) > 0 && (seed%3 == 0 || len(reference) > 80) {
			s.Pop()
			reference = reference[:len(reference)-1]
		} else {
			val := int(int32(uint32(seed >> 16)))
			if step%7 == 0 {
				val = -2147483648
			}
			if step%11 == 0 {
				val = 2147483647
			}
			if step%13 == 0 {
				val = 0
			}
			s.Push(val)
			reference = append(reference, val)
		}
		if len(reference) > 0 {
			low := reference[0]
			for _, v := range reference {
				if v < low {
					low = v
				}
			}
			eq(t, s.GetMin(), low)
			eq(t, s.Top(), reference[len(reference)-1])
		}
	}
	for len(reference) > 0 {
		s.Pop()
		reference = reference[:len(reference)-1]
	}
	s.Push(2147483647)
	s.Push(-2147483648)
	eq(t, s.Top(), -2147483648)
	eq(t, s.GetMin(), -2147483648)
	s.Pop()
	eq(t, s.Top(), 2147483647)
	eq(t, s.GetMin(), 2147483647)
	s.Pop()
	s.Push(-42)
	eq(t, s.Top(), -42)
	eq(t, s.GetMin(), -42)
	t.Logf("checks=%d", checks)
}
