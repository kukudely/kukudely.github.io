package solution

import (
	"leetcodepractice/internal/testutil"
	"testing"
)

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
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	t.Run("split", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		root := tree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
		if lowestCommonAncestor(root, root.Left, root.Right) != root {
			t.Fatal("incorrect LCA")
		}
	})
	t.Run("ancestor", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		root := tree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
		if lowestCommonAncestor(root, root.Left, root.Left.Right.Right) != root.Left {
			t.Fatal("incorrect LCA")
		}
	})
	t.Run("deep_same_subtree", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		root := tree(3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4)
		if lowestCommonAncestor(root, root.Left.Right.Left, root.Left.Right.Right) != root.Left.Right {
			t.Fatal("incorrect LCA")
		}
	})
	t.Run("root_target", func(t *testing.T) {
		defer testutil.SkipTODO(t)
		root := tree(1, 2)
		if lowestCommonAncestor(root, root, root.Left) != root {
			t.Fatal("incorrect LCA")
		}
	})
}
