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
func preorderValues(root *TreeNode) []int {
	values := []int{}
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		values = append(values, node.Val)
		walk(node.Left)
		walk(node.Right)
	}
	walk(root)
	return values
}
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	roots := []*TreeNode{tree(1, 2, 5, 3, 4, nil, 6), nil, tree(1), tree(1, 2, nil, 3), tree(1, nil, 2, nil, 3)}
	for i, root := range roots {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			want := preorderValues(root)
			nodes := map[*TreeNode]bool{}
			var visit func(*TreeNode)
			visit = func(node *TreeNode) {
				if node == nil {
					return
				}
				nodes[node] = true
				visit(node.Left)
				visit(node.Right)
			}
			visit(root)
			flatten(root)
			got := []int{}
			for node := root; node != nil; node = node.Right {
				if node.Left != nil {
					t.Fatal("left pointer retained")
				}
				if !nodes[node] {
					t.Fatal("new node or repeated node")
				}
				delete(nodes, node)
				got = append(got, node.Val)
			}
			eq(t, got, want)
			eq(t, len(nodes), 0)
		})
	}
}
