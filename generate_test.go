package stack

import "testing"

var expected = `package stack

import "errors"

type TestStack struct {
	s Stack
}

func (s *TestStack) pop() (Test, error) {
	var auto Test
	top, err := s.s.pop()
	if err != nil {
		return auto, err
	}
	topInstance, correct := top.(Test)
	if !correct {
		return auto, errors.New("Popped element has invalid type. Expected Test.")
	}
	return topInstance, nil
}

func (s *TestStack) push(val Test) {
	s.s.push(val)
	return
}
`

func TestGenerate(t *testing.T) {
	codeForTestType := GenerateStackType("Test")

	if codeForTestType != expected {
		t.Error("Generated code differs from expected")
		t.Error("Generated", codeForTestType)
		t.Error("Expected", expected)
	}

}

// TestWriteToFile is written for local testing only
/*
func TestWriteToFile(t *testing.T) {
	gopath := os.Getenv("GOPATH")
	writeLocation := gopath + "\\src\\github.com\\voutasaurus\\stack\\teststack.go"
	WriteStackTypeFile("Test", writeLocation)

	file, err := os.Open(writeLocation)
	if err != nil {
		t.Error(err.Error())
	}
	readBytes := make([]byte, 408)
	n, err := file.Read(readBytes)
	if err != nil {
		t.Error("Read bytes:", n, "Error:", err.Error())
	}

	s := bytes.NewBuffer(readBytes).String()
	if s != expected {
		t.Error("Expected:", expected, "Got:", s)
	}
}
*/
