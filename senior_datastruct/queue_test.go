package senior_datastruct

import (
	"fmt"
	"math/rand"
	"testing"
)

func TestQueue(t *testing.T) {

	queue := NewPriorityQueue(10)

	for i := 0; i < 10; i++ {
		queue.Enqueue(rand.Intn(50))
	}

	for !queue.IsEmpty() {
		val, _ := queue.Dequeue()
		fmt.Println(val)
	}
}
