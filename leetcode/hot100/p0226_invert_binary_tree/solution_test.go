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
		name       string
		root, want *TreeNode
	}{{"official", tree(4, 2, 7, 1, 3, 6, 9), tree(4, 7, 2, 9, 6, 3, 1)}, {"empty", nil, nil}, {"single", tree(1), tree(1)}, {"chain", tree(1, 2, nil, 3), tree(1, nil, 2, nil, 3)}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			before := treeKey(c.root)
			got := invertTree(c.root)
			if got != c.root {
				t.Fatal("root pointer changed")
			}
			eq(t, treeKey(got), treeKey(c.want))
			eq(t, treeKey(invertTree(got)), before)
		})
	}
}
