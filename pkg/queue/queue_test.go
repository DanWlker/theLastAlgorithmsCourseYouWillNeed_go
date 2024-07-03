package queue

import "testing"

func TestQueue(t *testing.T) {

	var list Queue

	list.Enqueue(5)
	list.Enqueue(7)
	list.Enqueue(9)

	if val := list.Deque(); val != 5 {
		t.Errorf("Expected %v, got %v", 5, val)
	}

	if val := list.Length(); val != 2 {
		t.Errorf("Expected %v, got %v", 2, val)
	}

	list.Enqueue(11)

	if val := list.Deque(); val != 7 {
		t.Errorf("Expected %v, got %v", 7, val)
	}

	if val := list.Deque(); val != 9 {
		t.Errorf("Expected %v, got %v", 9, val)
	}

	if val := list.Peek(); val != 11 {
		t.Errorf("Expected %v, got %v", 11, val)
	}

	if val := list.Deque(); val != 11 {
		t.Errorf("Expected %v, got %v", 11, val)
	}

	if val := list.Deque(); val != 0 {
		t.Errorf("Expected %v, got %v", 0, val)
	}

	if val := list.Length(); val != 0 {
		t.Errorf("Expected %v, got %v", 0, val)
	}

	list.Enqueue(69)

	if val := list.Peek(); val != 69 {
		t.Errorf("Expected %v, got %v", 69, val)
	}

	if val := list.Length(); val != 1 {
		t.Errorf("Expected %v, got %v", 1, val)
	}
}
