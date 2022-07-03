package set_test

import (
	"github.com/LarsNorlander/aoc-2021/lib/set"
	"testing"
)

func TestSetOf(t *testing.T) {
	setA := set.Of("A", "B")
	if !setA.Contains("A") || !setA.Contains("B") {
		t.Log("the set should have had A and B")
		t.Failed()
	}
}

func TestUnion(t *testing.T) {
	setA := set.Of("A", "B")
	setB := set.Of("C")
	setC := set.Union(setA, setB)
	if !setC.Contains("A") ||
		!setC.Contains("B") ||
		!setC.Contains("C") {
		t.Log("the set should have had A, B, and C")
		t.Failed()
	}
}

func TestSet_Append(t *testing.T) {
	setA := set.Of("A")
	setA.Add("B")
	if !setA.Contains("B") {
		t.Log("the set should have had B")
	}
}
