package day07

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// A type is Measurable if we can get its size
type Measurable interface {
	Size() int
}

// A Folder has a name, a list of children and a cached size
type Folder struct {
	Name     string
	Children [](Measurable)
	size     int
}

// The size of a Folder is the size of their contents
func (folder *Folder) Size() int {
	size := 0

	// Get the computed size if present
	if folder.size > 0 {
		return folder.size
	}

	// If not, compute it recursively
	for _, child := range folder.Children {
		size += (child).Size()
	}

	// Save the size for later
	folder.size = size

	return size
}

// A size has a name and a size
type File struct {
	Name string
	size int
}

// A file's size is constant
func (file File) Size() int {
	return file.size
}

// Parse a Folder recursively using a Scanner
func parseFolder(name string, scanner *bufio.Scanner) (Folder, *bufio.Scanner) {
	children := [](Measurable){}

	for scanner.Scan() {
		line := scanner.Text()

		if line == "$ cd .." {
			// Stop parsing once we return to the upper level
			break
		} else if line == "$ ls" {
			// Ignore the ls statements
			continue
		} else if line[:4] == "dir " {
			// Ignore the dir statements - they will be parsed on cd
			continue
		} else if line[:4] == "$ cd" {
			// Parse a folder once we cd in it
			var child Folder
			child, scanner = parseFolder(line[5:], scanner)
			children = append(children, &child)
		} else {
			// Parse a file and add them to the Folder's children
			values := strings.Split(line, " ")
			size, err := strconv.Atoi(values[0])
			if err != nil {
				panic(fmt.Sprintf("Error parsing %s", line))
			}

			child := File{
				size: size,
				Name: values[1],
			}

			children = append(children, child)
		}
	}

	// Generate the resulting folder with an invalid cached size
	folder := Folder{
		Name:     name,
		Children: children,
		size:     -1,
	}

	return folder, scanner
}

// Read the input as the root folder
func readInput(path string) Folder {
	var input Folder

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Discard the first line - it is "cd /"
		_ = scanner.Text()
		input, scanner = parseFolder("/", scanner)
	}

	return input
}
