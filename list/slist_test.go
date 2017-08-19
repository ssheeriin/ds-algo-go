package list

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func TestSLinkedList_AddToHead(t *testing.T) {
	c := new(SLinkedList)
	c.AddToHead(10)
	c.AddToHead(11)
	c.AddToHead(12)

	out := printListAndCaptureOut(c)

	excpected := "12  11  10  \n"
	if out != excpected {
		t.Errorf("got: %s, want: %s", out, excpected)
	}
}

func TestSLinkedList_Print(t *testing.T) {
	n3 := &slnode{data: 3, next: nil}
	n2 := &slnode{data: 2, next: n3}
	n1 := &slnode{data: 1, next: n2}
	c := new(SLinkedList)
	c.head = n1

	out := printListAndCaptureOut(c)

	excpected := "1  2  3  \n"
	if out != excpected {
		t.Errorf("got: %s, want: %s", out, excpected)
	}
}

func printListAndCaptureOut(c *SLinkedList) string {
	// capture std out
	old := os.Stdout
	// keep backup of the real stdout
	r, w, _ := os.Pipe()
	os.Stdout = w
	// do printing
	c.Print()
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
