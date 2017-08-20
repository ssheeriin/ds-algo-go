package tree

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestBSTree_Add(t *testing.T) {
	bst := new(BSTree)
	bst.Add(1)
	bst.Add(5)
	bst.Add(3)
	bst.Add(7)
	bst.Add(9)
	bst.Add(2)
	bst.Add(6)
	bst.Add(4)
	bst.Add(8)

	traverse := bst.InOrder
	out := executeFuncAndCaptureOut(traverse)

	expected := "123456789"
	if expected != out {
		t.Errorf("expected: %s, actual:%s", expected, out)
	}
}

type traverse func()

func executeFuncAndCaptureOut(t traverse) string {
	// capture std out
	old := os.Stdout
	// keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// do printing
	t()
	outC := make(chan string)
	// copy the output in a separate goroutine so printing can't block indefinitely
	go func() {
		var buf bytes.Buffer
		io.Copy(&buf, r)
		outC <- buf.String()
	}()
	// back to normal state
	w.Close()
	os.Stdout = old
	// restoring the real stdout
	out := <-outC
	return out
}
