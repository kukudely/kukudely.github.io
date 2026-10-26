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
	cases := []struct {
		name string
		root *TreeNode
		want []int
	}{{"official", tree(1, nil, 2, 3), []int{1, 3, 2}}, {"empty", nil, []int{}}, {"single", tree(-1), []int{-1}}, {"left_chain", tree(3, 2, nil, 1), []int{1, 2, 3}}, {"balanced", tree(4, 2, 6, 1, 3, 5, 7), []int{1, 2, 3, 4, 5, 6, 7}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			before := treeKey(c.root)
			eq(t, inorderTraversal(c.root), c.want)
			eq(t, treeKey(c.root), before)
		})
	}
}
