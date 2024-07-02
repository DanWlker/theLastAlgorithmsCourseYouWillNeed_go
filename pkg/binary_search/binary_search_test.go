package binarysearch

import "testing"

func TestBinarySearch(t *testing.T) {
	foo := []int{1, 3, 4, 69, 71, 81, 90, 99, 420, 1337, 69420}
	ex := []struct {
		target int
		result bool
	}{
		{69, true},
		{1336, false},
		{69420, true},
		{69421, false},
		{1, true},
		{0, false},
	}

	for _, exV := range ex {
		if res := BinarySearch(foo, exV.target); exV.result != res {
			t.Errorf("Expected %v, got %v", exV.result, res)
		}
	}

}
