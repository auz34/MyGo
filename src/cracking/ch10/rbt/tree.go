package rbt

type Node struct {
	Value int
	Left *Node
	Right *Node
	color bool // Black - false, Red - true
	size int
}

type RedBlackTree struct {
	root *Node
}

func NewRedBlackTree() *RedBlackTree {
	return &RedBlackTree{ nil }
}

func (self *RedBlackTree) Insert(value int) {
	self.root = insert(self.root, value)
	self.root.color = false // BLACK
}

func (self *RedBlackTree) GetRank(value int) int {
	return getRank(self.root, value)
}

func isRed(h *Node) bool {
	if h == nil {
		return false
	}

	return h.color
}

func size(h *Node) int {
	if h == nil {
		return 0
	}

	return h.size
}

func insert(h *Node, value int) *Node {
	if h == nil {
		return &Node{value, nil, nil, true, 1}
	}

	comp := h.Value - value
	if comp < 0 {
		h.Right = insert(h.Right, value)
	} else {
		h.Left = insert(h.Left, value)
	}

	// fix-up any right-leaning links
	if isRed(h.Right) && !isRed(h.Left) {
		h = rotateLeft(h)
	}

	if isRed(h.Left)  && isRed(h.Left.Left) {
		h = rotateRight(h)
	}

	if isRed(h.Left)  && isRed(h.Right) {
		flipColors(h);
	}

	h.size = size(h.Left) + size(h.Right) + 1

	return h
}

// make a left-leaning link lean to the right
func rotateRight(h *Node) *Node {
	x := h.Left
	h.Left = x.Right
	x.Right = h
	x.color = x.Right.color
	x.Right.color = true // RED
	x.size = h.size
	h.size = size(h.Left) + size(h.Right) + 1
	return x
}

func rotateLeft(h *Node) *Node {
	x := h.Right
	h.Right = x.Left
	x.Left = h
	x.color = x.Left.color
	x.Left.color = true //RED;
	x.size = h.size;
	h.size = size(h.Left) + size(h.Right) + 1
	return x;
}

func flipColors(h *Node) {
	h.color = !h.color
	h.Left.color = !h.Left.color
	h.Right.color = !h.Right.color
}

func getRank(h *Node, value int) int {
	if h == nil {
		return 0;
	}

	if h.Value == value {
		return size(h.Left)
	}

	comp := h.Value - value
	if comp > 0 {
		return getRank(h.Left, value)
	} else {
		return size(h.Left) + 1 + getRank(h.Right, value)
	}
}