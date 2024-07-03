package stack

import "testing"

func TestStack(t *testing.T) {
	var list Stack

	list.Push(5)
	list.Push(7)
	list.Push(9)

	if val := list.Pop(); val != 9 {
		t.Errorf("Expected %v, got %v", 9, val)
	}

	if val := list.Length(); val != 0 {
		t.Errorf("Expected %v, got %v", 0, val)
	}

	list.Push(11)

	if val := list.Pop(); val != 11 {
		t.Errorf("Expected %v, got %v", 11, val)
	}
	if val := list.Pop(); val != 7 {
		t.Errorf("Expected %v, got %v", 7, val)
	}
	if val := list.Peek(); val != 5 {
		t.Errorf("Expected %v, got %v", 5, val)
	}
	if val := list.Pop(); val != 5 {
		t.Errorf("Expected %v, got %v", 5, val)
	}
	if val := list.Pop(); val != 0 {
		t.Errorf("Expected %v, got %v", 0, val)
	}

	list.Push(69)

	if val := list.Peek(); val != 69 {
		t.Errorf("Expected %v, got %v", 69, val)
	}
	if val := list.Length(); val != 1 {
		t.Errorf("Expected %v, got %v", 1, val)
	}
}
