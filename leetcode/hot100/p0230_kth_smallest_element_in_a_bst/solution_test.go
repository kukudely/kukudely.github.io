package solution

import (
	"fmt"
	"leetcodepractice/internal/testutil"
	"reflect"
	"testing"
)

func eq(t *testing.T, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Fatalf("got=%v; want=%v", got, want)
	}
}
func tree(values ...any) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}
	root := &TreeNode{Val: values[0].(int)}
	q := []*TreeNode{root}
	index := 1
	for head := 0; head < len(q) && index < len(values); head++ {
		node := q[head]
		if values[index] != nil {
			node.Left = &TreeNode{Val: values[index].(int)}
			q = append(q, node.Left)
		}
		index++
		if index < len(values) {
			if values[index] != nil {
				node.Right = &TreeNode{Val: values[index].(int)}
				q = append(q, node.Right)
			}
			index++
		}
	}
	return root
}
func treeKey(root *TreeNode) string {
	seen := map[*TreeNode]bool{}
	var key func(*TreeNode) string
	key = func(node *TreeNode) string {
		if node == nil {
			return "#"
		}
		if seen[node] {
			return "CYCLE"
		}
		seen[node] = true
		return fmt.Sprintf("%d(%s)(%s)", node.Val, key(node.Left), key(node.Right))
	}
	return key(root)
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	t.Run("official", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		eq(t, kthSmallest(tree(3, 1, 4, nil, 2), 1), 1)
	})
	for k := 1; k <= 7; k++ {
		t.Run(fmt.Sprintf("rank_%d_and_restore", k), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			root := tree(4, 2, 6, 1, 3, 5, 7)
			before := treeKey(root)
			eq(t, kthSmallest(root, k), k)
			eq(t, treeKey(root), before)
		})
	}
	t.Run("single", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		eq(t, kthSmallest(tree(0), 1), 0)
	})
	t.Run("left_chain", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		eq(t, kthSmallest(tree(3, 2, nil, 1), 3), 3)
	})
}
