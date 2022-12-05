package dsa

// A stack of integers is represented internally with a slice
type IntStack struct {
	rel []int
}

// Create a new IntStack and initialize all fields
func NewIntStack() IntStack {
	return IntStack{[]int{}}
}

// Add a new element to the IntStack
func (stack *IntStack) Push(elem int) {
	stack.rel = append(stack.rel, elem)
}

// Pop the last element pushed to the IntStack
func (stack *IntStack) Pop() int {
	result := stack.rel[len(stack.rel)-1]
	stack.rel = stack.rel[0 : len(stack.rel)-1]
	return result
}

// Return the top element of the stack but do not remove it
func (stack IntStack) Peek() int {
	return stack.rel[len(stack.rel)-1]
}

// Return a slice of the values of an IntStack
func (stack IntStack) Values() []int {
	return stack.rel
}
