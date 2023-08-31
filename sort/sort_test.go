package sort

import (
	"github.com/stretchr/testify/assert"
	"log"
	"math/rand"
	"testing"
	"time"
)

var (
	arr    []int
	length int
)

func init() {
	//rand.Seed(time.Now().Unix())
	rand.NewSource(time.Now().Unix())
	length = rand.Intn(26)
	arr = make([]int, length)
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Intn(100)
	}
	log.Printf("arr init success: \n %v\n", arr)
}

func validateMax(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] < arr[i+1] {
			return false
		}
	}

	return true
}

func TestSelectSort(t *testing.T) {

	SelectSort(arr)
	log.Printf("TestSelectSort res: \n %v\n", arr)
	valid := validateMax(arr)
	assert.Equal(t, true, valid)
}

func TestQuickSort(t *testing.T) {
	a := []int{83, 90, 29, 82, 49, 21, 68, 9, 71, 6, 1, 85, 89}
	QuickSort(a)
}
