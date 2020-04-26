package priorityqueue

type LeftishTree *Node

type Node struct {
	Key         int
	Dist        int
	Left, Right LeftishTree
	Info        interface{}
}

func Merge(a, b LeftishTree) LeftishTree {
	if a == nil {
		return b
	}

	if b == nil {
		return a
	}

	if a.Key > b.Key {
		a, b = b, a
	} // now a.key <= b.key

	a.Right = Merge(a.Right, b)

	if a.Left == nil && a.Right != nil ||
		a.Left != nil && a.Right != nil &&
			a.Right.Dist > a.Left.Dist {
		a.Right, a.Left = a.Left, a.Right
	}

	if a.Right == nil {
		a.Dist = 1
	} else {
		a.Dist = a.Right.Dist + 1
	}

	return a
}

func Equal(a, b LeftishTree) bool {
	if a == nil {
		return b == nil
	}

	if b == nil {
		return a == nil
	}

	return a.Key == b.Key &&
		a.Dist == b.Dist &&
		Equal(a.Left, b.Left) &&
		Equal(a.Right, b.Right) &&
		a.Info == b.Info
}
