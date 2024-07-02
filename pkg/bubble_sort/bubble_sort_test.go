package bubblesort

import "testing"

func TestBubbleSort(t *testing.T) {
	arr := []int{9, 3, 7, 4, 69, 420, 42}
	ex := []int{3, 4, 7, 9, 42, 69, 420}

	BubbleSort(&arr)
	for i, exV := range ex {
		if exV != arr[i] {
			t.Errorf("Expected %v, got %v", ex, arr)
			break
		}
	}
}
