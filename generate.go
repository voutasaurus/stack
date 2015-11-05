package stack

import (
	"fmt"
	"os"
	"strings"
)

var template = `package stack

import "errors"

type xkcdStack struct {
	s Stack
}

func (s *xkcdStack) pop() (xkcd, error) {
	var auto xkcd
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(xkcd)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected xkcd.")
	}
	return topInstance, nil
}

func (s *xkcdStack) push(val xkcd) {
	s.s.push(val)
	return
}
`

// GenerateStackType generates the code for a stack of elements of type typeName using a standard template
func GenerateStackType(typeName string) string {
	return strings.Replace(template, "xkcd", typeName, -1)
}

// WriteStackTypeFile writes the code for a stack of elements of type typeName at the filesystem location given by path
func WriteStackTypeFile(typeName string, path string) {
	toWrite := GenerateStackType(typeName)

	file, err := os.Create(path)
	if err != nil {
		fmt.Println("File could not be created at", path, "Error:", err)
	}

	n, err := file.WriteString(toWrite)
	if err != nil {
		fmt.Println("Bytes written", n, "Error:", err)
	}

	return
}
