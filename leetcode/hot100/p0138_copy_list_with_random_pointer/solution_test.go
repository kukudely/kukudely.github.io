package solution

import (
	"leetcodepractice/internal/testutil"
	"math/rand"
	"testing"
)

func TestCases(t *testing.T) {
	defer testutil.SkipTODO(t)
	if copyRandomList(nil) != nil {
		t.Fatal("nil copy")
	}
	rng := rand.New(rand.NewSource(138))
	for n := 1; n <= 25; n++ {
		for trial := 0; trial < 15; trial++ {
			original := make([]*Node, n)
			randomIndices := make([]int, n)
			for i := range original {
				original[i] = &Node{Val: rng.Intn(5)}
			}
			for i, p := range original {
				if i+1 < n {
					p.Next = original[i+1]
				}
				j := rng.Intn(n+1) - 1
				randomIndices[i] = j
				if j >= 0 {
					p.Random = original[j]
				}
			}
			result := copyRandomList(original[0])
			clones := []*Node{}
			visited := map[*Node]bool{}
			oldSet := map[*Node]bool{}
			for _, p := range original {
				oldSet[p] = true
			}
			for p := result; p != nil; p = p.Next {
				if oldSet[p] || visited[p] {
					t.Fatal("copy alias or cycle")
				}
				visited[p] = true
				clones = append(clones, p)
			}
			if len(clones) != n {
				t.Fatal("wrong copy length")
			}
			for i, p := range original {
				var next *Node
				if i+1 < n {
					next = original[i+1]
				}
				if p.Next != next {
					t.Fatal("original next not restored")
				}
				j := randomIndices[i]
				var oldRandom, newRandom *Node
				if j >= 0 {
					oldRandom = original[j]
					newRandom = clones[j]
				}
				if p.Random != oldRandom || clones[i].Random != newRandom || clones[i].Val != p.Val {
					t.Fatal("random mapping/value mismatch")
				}
			}
		}
	}
}
