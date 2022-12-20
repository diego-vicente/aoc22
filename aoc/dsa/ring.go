package dsa

// A Node holds a value and forward and backward pointers
type Node[T comparable] struct {
	Value T
	Prev  *Node[T]
	Next  *Node[T]
}

// A ring is a singly linked list with extra logic
type Ring[T comparable] struct {
	start *Node[T]
	end   *Node[T]
	rel   []*Node[T]
	Size  int
}

// Create a new empty Ring
func NewRing[T comparable]() Ring[T] {
	return Ring[T]{rel: []*Node[T]{}, Size: 0}
}

// Add a new item to the end of the Ring
func (ring *Ring[T]) Append(elem T) {
	node := Node[T]{Value: elem}

	ring.rel = append(ring.rel, &node)
	ring.Size += 1

	if ring.start == nil {
		ring.start = &node
	}

	if old := ring.end; old == nil {
		ring.end = &node
	} else {
		node.Prev = old
		old.Next = &node
		ring.end = &node
	}
}

// Get the node at a given index of the ring
func (ring *Ring[T]) GetNode(idx int) *Node[T] {
	if ring.Size == 0 {
		return nil
	}

	idx = Mod(idx, ring.Size)
	ptr := ring.start
	for i := 0; i < idx; i++ {
		ptr = ptr.Next
	}

	return ptr
}

// Get the value at a given index of the ring
func (ring *Ring[T]) GetValue(idx int) T {
	return ring.GetNode(idx).Value
}

// Remove an item at a given index and return the node
func (ring *Ring[T]) pop(idx int) *Node[T] {
	node := ring.GetNode(idx)

	if ring.start == node {
		ring.start = node.Next
		ring.start.Prev = nil
	} else if ring.end == node {
		ring.end = node.Prev
		ring.end.Next = nil
	} else {
		prev := node.Prev
		next := node.Next

		prev.Next = next
		next.Prev = prev
	}

	node.Prev = nil
	node.Next = nil

	ring.Size--
	return node
}

// Insert an node just after a given index
func (ring *Ring[T]) insertAfter(node, prev *Node[T]) {
	if prev == ring.end {
		ring.end = node
		node.Prev = prev
		prev.Next = node
	} else {
		next := prev.Next

		node.Next = next
		node.Prev = prev

		next.Prev = node
		prev.Next = node
	}

	ring.Size++
}

// Insert a node at a certain index
func (ring *Ring[T]) insertAt(node *Node[T], idx int) {
	if idx == 0 {
		// Wrap around the list
		ring.insertAfter(node, ring.end)
	} else if prev := ring.GetNode(idx).Prev; prev != nil {
		ring.insertAfter(node, prev)
	} else {
		ring.insertAfter(node, ring.end)
	}
}

// Move a given index a certain number of steps forward
func (ring *Ring[T]) MoveIdx(idx, delta int) {
	dest := idx + delta

	// Pop and insert again
	node := ring.pop(idx)
	ring.insertAt(node, dest)
}

// Move a given node a certain number of elements forward
func (ring *Ring[T]) MoveNode(node *Node[T], delta int) {
	if delta == 0 {
		return
	}

	// Find the current node index
	idx := 0
	ptr := ring.start
	for i := 0; i < ring.Size && ptr != nil; i++ {
		if node == ptr {
			idx = i
			break
		} else {
			ptr = ptr.Next
		}
	}

	ring.MoveIdx(idx, delta)
}

// Get the slice of values in the current order
func (ring Ring[T]) Values() []T {
	values := []T{}

	ptr := ring.start
	for ptr != nil {
		values = append(values, ptr.Value)
		ptr = ptr.Next
	}

	return values
}

// Get the slice of nodes in the insertion order
func (ring Ring[T]) OriginalNodes() []*Node[T] {
	return ring.rel
}

// Get the list of values in the insertion order
func (ring Ring[T]) OriginalValues() []T {
	values := []T{}

	for _, node := range ring.rel {
		values = append(values, node.Value)
	}

	return values
}

// Find the node index holding a given value
func (ring Ring[T]) FindIndex(elem T) int {
	idx := 0
	ptr := ring.start
	for ptr != nil {
		if ptr.Value == elem {
			return idx
		} else {
			idx++
			ptr = ptr.Next
		}
	}
	return -1
}
