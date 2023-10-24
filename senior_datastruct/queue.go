package senior_datastruct

import "errors"

// Queue
// left = 2 * i + 1
// right = 2 * + 2
//
// parent = (i - 1) / 2
type Queue interface {
	Enqueue(int)
	Dequeue() (int, error)
	GetFront() (int, error)
	GetSize() int
	IsEmpty() bool
}

type MaxHeap struct {
	Data []int
}

func NewPriorityQueue(capacity int) Queue {
	return &MaxHeap{
		Data: make([]int, 0, capacity),
	}
}

func (m *MaxHeap) Enqueue(val int) {
	// 元素添加到末尾
	m.Data = append(m.Data, val)
	m.siftUp(len(m.Data) - 1)
}

func (m *MaxHeap) Dequeue() (int, error) {
	if len(m.Data) == 0 {
		return 0, errors.New("queue is empty")
	}
	res := m.Data[0]
	swap(0, len(m.Data)-1, m.Data)
	m.Data = m.Data[0 : len(m.Data)-1]

	// 第一个元素下沉
	m.siftDown(0)

	return res, nil
}

func (m *MaxHeap) GetFront() (int, error) {
	if len(m.Data) == 0 {
		return 0, errors.New("queue is empty")
	}
	return m.Data[0], nil
}

func (m *MaxHeap) GetSize() int {
	return len(m.Data)
}

func (m *MaxHeap) IsEmpty() bool {
	return len(m.Data) == 0
}

// private method
func (m *MaxHeap) parent(index int) int {

	return (index - 1) / 2
}

func (m *MaxHeap) leftChild(index int) int {

	return index*2 + 1
}

func (m *MaxHeap) rightChild(index int) int {

	return index*2 + 2
}

func (m *MaxHeap) siftUp(index int) {
	// parent value < k value
	for index > 0 && m.Data[m.parent(index)] < m.Data[index] {
		swap(index, m.parent(index), m.Data)
		index = m.parent(index)
	}
}

func (m *MaxHeap) siftDown(index int) {
	for m.leftChild(index) < len(m.Data) {
		j := m.leftChild(index)
		if j+1 < len(m.Data) &&
			m.Data[j+1] > m.Data[j] {
			// 右节点索引
			j = j + 1
		}
		// 当前节点满足大根堆
		if m.Data[index] > m.Data[j] {
			break
		} else {
			swap(index, j, m.Data)
			index = j
		}
	}
}

func swap(i, j int, arr []int) {
	tmp := arr[i]
	arr[i] = arr[j]
	arr[j] = tmp
}
