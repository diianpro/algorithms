package traverse

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type Node struct {
	Val      int
	Children []*Node
}

type stack struct {
	st []*Node
}

func (s *stack) pop() *Node {
	node := s.st[0]
	s.st = s.st[1:]
	return node
}

func (s *stack) push(node *Node) {
	s.st = append([]*Node{node}, s.st...)

}

func traverse(node *Node) []int {
	res := make([]int, 0)
	s := stack{
		st: make([]*Node, 0),
	}
	s.push(node)

	for len(s.st) != 0 {
		n := s.pop()
		res = append(res, n.Val)
		for i := len(n.Children) - 1; i >= 0; i-- {
			s.push(n.Children[i])
		}
	}
	return res
}

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	return traverse(root)
}

func TestIsHappy(t *testing.T) {
	cases := []struct {
		name   string
		root   *Node
		result []int
	}{
		{
			name: "1",
			root: &Node{
				Val: 1,
				Children: []*Node{
					&Node{
						Val: 3,
						Children: []*Node{
							&Node{
								Val: 5,
								Children: []*Node{
									&Node{
										Val:      7,
										Children: []*Node{},
									},
								},
							},
							&Node{
								Val:      6,
								Children: []*Node{},
							},
						},
					},
					&Node{
						Val:      2,
						Children: []*Node{},
					},
					&Node{
						Val:      4,
						Children: []*Node{},
					},
				},
			},
			result: []int{1, 3, 5, 7, 6, 2, 4},
		},
	}
	for _, cc := range cases {
		tt := t
		tt.Run(cc.name, func(t *testing.T) {
			require.Equal(t, cc.result, preorder(cc.root))
		})
	}
}
