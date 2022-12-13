package day13

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// A Packet implements a Compare method to other Packets
type Packet interface {
	Compare(Packet) Comparison
}

// A ScalarPacket holds a single integer value
type ScalarPacket struct {
	Value int
}

// A ListPacket holds a list of Packets
type ListPacket struct {
	Value []Packet
}

// Convert a ScalarPacket to a ListPacket holding a single value
func (packet ScalarPacket) ToList() ListPacket {
	return ListPacket{[]Packet{packet}}
}

// A PacketPair has a first and a second Packet
type PacketPair struct {
	First  Packet
	Second Packet
}

// Parse a ScalarPacket as a string representation of an integer
func parseScalarPacket(token string) ScalarPacket {
	scalar, err := strconv.Atoi(token)
	if err != nil {
		panic(fmt.Sprintf("Error parsing token: %s", token))
	}

	return ScalarPacket{Value: scalar}
}

// Parse a Packet recursively from a given pointer onwards
func parsePacket(line string, ptr int) (Packet, int) {
	packet := ListPacket{[]Packet{}}
	token := ""

	for i := ptr; i < len(line); i++ {
		switch rune(line[i]) {
		case ']':
			if token != "" {
				packet.Value = append(packet.Value, parseScalarPacket(token))
			}

			return packet, i
		case '[':
			child, end := parsePacket(line, i+1)
			packet.Value = append(packet.Value, child)
			i = end
		case ',':
			if token != "" {
				packet.Value = append(packet.Value, parseScalarPacket(token))
				token = ""
			}
		default:
			token += string(line[i])
		}
	}

	panic("Error, imabalanced brackets")
}

// Read the first part input as a list of PacketPairs
func readFirstPartInput(path string) []PacketPair {
	var input []PacketPair
	var current []Packet

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			input = append(input, PacketPair{
				First:  current[0],
				Second: current[1],
			})

			current = []Packet{}
		} else {
			packet, _ := parsePacket(line, 1)
			current = append(current, packet)
		}
	}

	if len(current) > 0 {
		input = append(input, PacketPair{
			First:  current[0],
			Second: current[1],
		})
	}

	return input
}

// Read the second part input as a list of Packets
func readSecondPartInput(path string) []Packet {
	var input []Packet

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			continue
		} else {
			packet, _ := parsePacket(line, 1)
			input = append(input, packet)
		}
	}

	return input
}
