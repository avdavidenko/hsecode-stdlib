package tree

import "sort"
import "strconv"
import "errors"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func NewBST(elements []int) *Tree {
	h := make(map[int]struct{})

	unique := make([]int, 0)

	for _, e := range elements {
		_, ok := h[e]
		if !ok {
			h[e] = struct{}{}
			unique = append(unique, e)
		}
	}

	sort.Ints(unique)

	return fromSorted(unique)
}

func (T *Tree) IsSym() bool {
	return isMirror(T, T)
}

func (T *Tree) InOrder(visit func(node *Tree)) {
	stack := make([](*Tree), 0)
	curr := T

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		curr = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		visit(curr)

		curr = curr.Right
	}
}

func fromSorted(unique []int) *Tree {
	if len(unique) == 0 {
		return nil
	}

	if len(unique) == 1 {
		return &Tree{Value: unique[0], Left: nil, Right: nil}
	}

	med := len(unique) / 2
	return &Tree{Value: unique[med], Left: fromSorted(unique[:med]), Right: fromSorted(unique[med+1:])}
}

func (T *Tree) noLeft() (*Tree, *Tree) {
	end := T
	if T.Right != nil {
		T.Right, end = T.Right.noLeft()
	}

	if T.Left == nil {
		return T, end
	}

	root, toInsert := T.Left.noLeft()
	T.Left = nil
	toInsert.Right = T

	return root, end
}

func (T *Tree) NoLeft() *Tree {
	root, _ := T.noLeft()
	return root
}

func isMirror(T1, T2 *Tree) bool {
	if T1 == nil && T2 == nil {
		return true
	}

	if T1 == nil || T2 == nil {
		return false
	}

	if T1.Value != T2.Value {
		return false
	}

	return isMirror(T1.Left, T2.Right) && isMirror(T1.Right, T2.Left)
}

func (T *Tree) Encode() []string {
	nodeCount := 0
	T.InOrder(func(node *Tree) { nodeCount++ })
	if nodeCount <= 0 {
		return make([]string, 0)
	}

	nodes := make([](*Tree), 1, nodeCount)
	nodes[0] = T
	nodeCount--
	for i := 0; i < len(nodes) && nodeCount > 0; i++ {
		if nodes[i] == nil {
			nodes = append(nodes, nil)
			nodes = append(nodes, nil)
			continue
		}

		if nodes[i].Left != nil {
			nodes = append(nodes, nodes[i].Left)
			nodeCount--
		} else {
			nodes = append(nodes, nil)
		}

		if nodeCount <= 0 {
			continue
		}

		if nodes[i].Right != nil {
			nodes = append(nodes, nodes[i].Right)
			nodeCount--
		} else {
			nodes = append(nodes, nil)
		}
	}

	result := make([]string, 0, len(nodes))
	for _, v := range nodes {
		if v == nil {
			result = append(result, "nil")
		} else {
			result = append(result, strconv.Itoa(v.Value))
		}

	}
	return result
}

func Decode(data []string) (*Tree, error) {
	if len(data) == 0 {
		return nil, nil
	}

	nodes := make([](*Tree), 1, len(data))
	if data[0] == "nil" {
		return nil, errors.New("Wrong format")
	}
	value, err := strconv.Atoi(data[0])
	if err != nil {
		return nil, err
	}
	nodes[0] = &Tree{Value: value}
	dataCounter := 1
	for i := 0; i < len(nodes) && dataCounter < len(data); i++ {
		for j := 0; j < 2 && dataCounter < len(data); j++ {
			if nodes[i] == nil {
				dataCounter++
				continue
			}

			if data[dataCounter] == "nil" {
				nodes = append(nodes, nil)
				dataCounter++
				continue
			}

			value, err := strconv.Atoi(data[dataCounter])
			if err != nil {
				return nil, err
			}

			newNode := &Tree{Value: value}
			if j == 0 {
				nodes[i].Left = newNode
			} else {
				nodes[i].Right = newNode
			}
			nodes = append(nodes, newNode)
			dataCounter++
		}
	}

	return nodes[0], nil
}
