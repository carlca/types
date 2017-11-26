package stacks

import (
	"fmt"
	"testing"
)

func TestStack(t *testing.T) {
	var test Stack
	test.Push("1")
	test.Push("2")
	if test.Peek() != "2" {
		t.Error("test.Peek() expected 2, got ", test.Peek())
	}
	if test.Size() != 2 {
		t.Error("test.Size() expected 2, got ", test.Size())
	}
	test.Pop()
	if test.Peek() != "1" {
		t.Error("test.Pop() expected 1, got ", test.Peek())
	}
	if test.Empty() {
		t.Error("test.Empty() expected false, got ", test.Empty())
	}
	fmt.Println([]string(test))
}
