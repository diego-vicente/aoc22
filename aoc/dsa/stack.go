package dsa

// A stack is represented internally with a slice
type Stack[T any] struct {
	rel []T
}

// Create a new Stack and initialize all fields
func NewStack[T any]() Stack[T] {
	return Stack[T]{[]T{}}
}

// Add a new element to the Stack
func (stack *Stack[T]) Push(elem T) {
	stack.rel = append(stack.rel, elem)
}

// Pop the last element pushed to the Stack
func (stack *Stack[T]) Pop() T {
	result := stack.rel[len(stack.rel)-1]
	stack.rel = stack.rel[0 : len(stack.rel)-1]
	return result
}

// Return the top element of the Stack but do not remove it
func (stack Stack[T]) Peek() T {
	return stack.rel[len(stack.rel)-1]
}

// Return a slice of the values of a Stack
func (stack Stack[T]) Values() []T {
	return stack.rel
}
