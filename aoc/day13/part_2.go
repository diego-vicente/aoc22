package day13

import (
	"reflect"
	"sort"
)

// Solve the second part by sorting the input and finding the control packets
func solveSecondPart(path string) int {
	input := readSecondPartInput(path)

	// Hard-coded list of control packets
	control := []Packet{
		ListPacket{[]Packet{ListPacket{[]Packet{ScalarPacket{2}}}}},
		ListPacket{[]Packet{ListPacket{[]Packet{ScalarPacket{6}}}}},
	}

	input = append(input, control...)

	sort.Slice(input, func(i, j int) bool {
		return input[i].Compare(input[j]) == Less
	})

	result := 1
	for i, packet := range input {
		if reflect.DeepEqual(packet, control[0]) ||
			reflect.DeepEqual(packet, control[1]) {
			result *= i + 1
		}
	}

	return result
}
