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
func inorderValues(root *TreeNode) []int {
	values := []int{}
	var walk func(*TreeNode)
	walk = func(node *TreeNode) {
		if node == nil {
			return
		}
		walk(node.Left)
		values = append(values, node.Val)
		walk(node.Right)
	}
	walk(root)
	return values
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
	cases := []struct{ pre, in []int }{{[]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7}}, {[]int{}, []int{}}, {[]int{-1}, []int{-1}}, {[]int{3, 2, 1}, []int{1, 2, 3}}, {[]int{1, 2, 3}, []int{1, 2, 3}}, {[]int{4, 2, 1, 3, 6, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}}}
	for i, c := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			root := buildTree(c.pre, c.in)
			eq(t, preorderValues(root), c.pre)
			eq(t, inorderValues(root), c.in)
		})
	}
}
