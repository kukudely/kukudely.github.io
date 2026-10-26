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
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := [][]int{{-10, -3, 0, 5, 9}, {}, {1}, {1, 3}, {-2147483648, 0, 2147483647}}
	long := make([]int, 100)
	for i := range long {
		long[i] = i
	}
	cases = append(cases, long)
	for i, values := range cases {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			defer testutil.SkipTODO(t)
			root := sortedArrayToBST(values)
			eq(t, inorderValues(root), append([]int{}, values...))
			var height func(*TreeNode) int
			height = func(node *TreeNode) int {
				if node == nil {
					return 0
				}
				a, b := height(node.Left), height(node.Right)
				if a-b > 1 || b-a > 1 {
					t.Fatal("not height balanced")
				}
				return 1 + max(a, b)
			}
			height(root)
		})
	}
}
