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
	chain := &TreeNode{Val: 1000000000}
	node := chain
	for i := 1; i < 5; i++ {
		node.Left = &TreeNode{Val: 1000000000}
		node = node.Left
	}
	cases := []struct {
		name         string
		root         *TreeNode
		target, want int
	}{{"official", tree(10, 5, -3, 3, 2, nil, 11, 3, -2, nil, 1), 8, 3}, {"empty", nil, 0, 0}, {"zeros_branch_rollback", tree(0, 0, 0), 0, 5}, {"negative", tree(-1, -1), -1, 2}, {"cross_branch", tree(1, -1, -1), 0, 2}, {"int64_sum", chain, -294967296, 0}, {"single", tree(8), 8, 1}}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			defer testutil.SkipTODO(t)
			eq(t, pathSum(c.root, c.target), c.want)
		})
	}
}
