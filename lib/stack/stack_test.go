package stack_test

import (
	"github.com/LarsNorlander/aoc-2021/lib/stack"
	"testing"
)

func TestStack(t *testing.T) {
	stackA := stack.Empty[string]()
	stackA.Push("A")
	stackA.Push("B")
	if stackA.Peek() != "B" {
		t.Fatal("value should have been B")
	}
	valB, err := stackA.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if valB != "B" {
		t.Fatal("value should have been B")
	}
	valA, err := stackA.Pop()
	if err != nil {
		t.Fatal(err)
	}
	if valA != "A" {
		t.Fatal("value should have been A")
	}
	_, err = stackA.Pop()
	if err == nil {
		t.Fatal("err should have been set")
	}
}
