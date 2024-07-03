package twocrystalballs

import (
	"math/rand/v2"
	"testing"
)

func generateList[T any](size int, generator func(i int) T) []T {
	res := make([]T, size)
	for i := 0; i < size; i++ {
		res[i] = generator(i)
	}
	return res
}

func TestTwoCrystalBalls(t *testing.T) {
	for i := 0; i <= 100; i++ {
		idx := rand.IntN(10000)
		data := generateList(
			10000,
			func(j int) bool {
				return j >= idx
			},
		)

		if val := TwoCrystalBalls(data); val != idx {
			t.Errorf("Expected %v, got %v", idx, val)
		}
	}

	if val := TwoCrystalBalls(generateList(821, func(i int) bool { return i == 820 })); val != 820 {
		t.Errorf("Expected %v, got %v", val, 820)
	}

	if val := TwoCrystalBalls(generateList(821, func(i int) bool { return false })); val != -1 {
		t.Errorf("Expected %v, got %v", val, -1)
	}
}
