package dsa

// A set of integers is represented by a map of integers
type IntSet struct {
	rel map[int]bool
}

// Create a new IntSet and initialize all fields
func NewIntSet() IntSet {
	return IntSet{map[int]bool{}}
}

// Add a new element to the IntSet
func (set IntSet) Add(elem int) {
	set.rel[elem] = true
}

// Remove an element from an IntSet
func (set IntSet) Remove(elem int) {
	delete(set.rel, elem)
}

// Return a slice of the values fo an IntSet
func (set IntSet) Values() []int {
	var values []int

	for value := range set.rel {
		values = append(values, value)
	}

	return values
}

// Return true if the IntSet contains the element
func (set IntSet) Contains(elem int) bool {
	_, ok := set.rel[elem]
	return ok
}

// Return the union of two given IntSets
func (set IntSet) Union(other IntSet) IntSet {
	union := NewIntSet()

	for elem := range set.rel {
		union.Add(elem)
	}

	for elem := range other.rel {
		union.Add(elem)
	}

	return union
}

// Return the intersection of two given IntSets
func (set IntSet) Intersection(other IntSet) IntSet {
	intersect := NewIntSet()

	for elem := range set.rel {
		if other.Contains(elem) {
			intersect.Add(elem)
		}
	}

	return intersect
}
