package day13

// Create a comparison type since Go does not accept custom implementations of
// `comparable`
type Comparison = int8

const (
	Less    Comparison = -1
	Equal   Comparison = 0
	Greater Comparison = 1
)

// Reverse a Comparison value
func Not(c Comparison) Comparison {
	switch c {
	case Equal:
		return Equal
	case Greater:
		return Less
	case Less:
		return Greater
	default:
		panic("Impossible value")
	}
}

// Compare a ScalarPacket with another packet
func (p1 ScalarPacket) Compare(p2 Packet) Comparison {
	switch p2 := p2.(type) {
	case ScalarPacket:
		if p1.Value == p2.Value {
			return Equal
		} else if p1.Value > p2.Value {
			return Greater
		} else {
			return Less
		}
	case ListPacket:
		return p1.ToList().Compare(p2)
	default:
		panic("Impossible comparison")
	}
}

// Compare a ListPacket with another Packet
func (p1 ListPacket) Compare(p2 Packet) Comparison {
	switch p2 := p2.(type) {
	case ScalarPacket:
		// Use the logic defined in the ScalarPacket implementation
		return Not(p2.Compare(p1))
	case ListPacket:
		for i := range p1.Value {
			// If other runs out of items, this one is Greater
			if i >= len(p2.Value) {
				return Greater
			}

			// Compare both lists element-wise
			switch c := p1.Value[i].Compare(p2.Value[i]); c {
			case Equal:
				continue
			default:
				return c
			}
		}

		// If this ran out of values first, then the other list is Greater
		if len(p1.Value) < len(p2.Value) {
			return Less
		} else {
			// Otherwise, both lists are identical
			return Equal
		}
	default:
		panic("Impossible comparison")
	}
}

// Solve the first part by counting the indices of sorted pairs
func solveFirstPart(path string) int {
	input := readFirstPartInput(path)

	result := 0
	for i, pair := range input {
		if pair.First.Compare(pair.Second) == Less {
			result += i + 1
		}
	}

	return result
}
