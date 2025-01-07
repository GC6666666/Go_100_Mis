package tests

import (
	"G100/5-Avoiding_interface_pollution/copier"
	"bytes"
	"fmt"
	"strings"
	"testing"
)

func TestCopySourceToDest(t *testing.T) {
	const input = "foo"
	source := strings.NewReader(input)
	dest := bytes.NewBuffer(make([]byte, 0))

	err := copier.CopySourceToDest(source, dest)
	if err != nil {
		t.Fatalf("copySourceToDest returned an error: %v", err)
	}

	got := dest.String()
	if got != input {
		t.Errorf("expected: %s, got %s", input, got)
	} else {
		fmt.Printf("Test passed. Destination content: %s\n", got)
	}

}
