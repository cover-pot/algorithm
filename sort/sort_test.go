package sort

import (
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

var (
	arr    []int
	length int
	s      *Sort
)

func init() {
	log.SetFormatter(&log.TextFormatter{})
	s = new(Sort)
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

func validateMin(arr []int) bool {
	for i := 0; i < len(arr)-1; i++ {
		if arr[i] > arr[i+1] {
			return false
		}
	}

	return true
}

func TestSelectSort(t *testing.T) {

	s.SelectSort(arr)
	log.Printf("TestSelectSort res: \n %v\n", arr)
	valid := validateMax(arr)
	assert.Equal(t, true, valid)
}

func TestQuickSort(t *testing.T) {
	s.QuickSort(arr)
	log.Printf("TestQuickSort res: \n %v\n", arr)
	valid := validateMin(arr)
	assert.Equal(t, true, valid)
}

func TestBubbleSort(t *testing.T) {
	s.BubbleSort(arr)
	log.Printf("TestBubbleSort res: \n %v\n", arr)
	valid := validateMax(arr)
	assert.Equal(t, true, valid)
}

func TestInsertionSort(t *testing.T) {
	s.InsertionSort(arr)
	log.Printf("TestInsertionSort res: \n %v\n", arr)
	valid := validateMin(arr)
	assert.Equal(t, true, valid)
}

func TestHeapSort(t *testing.T) {
	s.HeapSort(arr)
	log.Printf("TestHeapSort res: \n %v\n", arr)
	valid := validateMax(arr)
	assert.Equal(t, true, valid)
}

func TestMergeSort(t *testing.T) {
	s.MergeSort(arr)
	log.Printf("TestMergeSort res: \n %v\n", arr)
	valid := validateMin(arr)
	assert.Equal(t, true, valid)
}
