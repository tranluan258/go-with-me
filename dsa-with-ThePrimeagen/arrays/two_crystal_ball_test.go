package arrays

import (
	"math/rand"
	"testing"
	"time"
)

func TestTwoCrystalBall(t *testing.T) {
	rand.Seed(time.Now().UnixNano())

	idx := rand.Intn(10000)

	data := make([]bool, 10000)

	for i := idx; i < 10000; i++ {
		data[i] = true
	}

	index := TwoCrystalBall(data)

	if index != idx {
		t.Fatalf("Expected %d rec %d", idx, index)
	}
}
