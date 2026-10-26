package solution

import (
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
func TestExamplesAndEdges(t *testing.T) {
	defer testutil.SkipTODO(t)
	cases := []struct {
		name string
		root *TreeNode
		want []int
	}{{"official", tree(1, 2, 3, nil, 5, nil, 4), []int{1, 3, 4}}, {"left_deeper", tree(1, 2, 3, 4, nil, nil, nil, 5), []int{1, 3, 4, 5}}, {"empty", nil, []int{}}, {"single", tree(1), []int{1}}, {"left_chain", tree(1, 2, nil, 3), []int{1, 2, 3}}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, rightSideView(c.root), c.want)
		})
	}
}
