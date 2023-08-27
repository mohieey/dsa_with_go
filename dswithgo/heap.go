package dswithgo

type MaxHeap struct {
	contents []int
}

// has o(log n) or o(h) complexity
func (m *MaxHeap) Insert(num int) {
	m.contents = append(m.contents, num)
	m.maxHeapifyUp()
}

// has o(log n) or o(h) complexity
func (m *MaxHeap) Extract() int {
	if len(m.contents) == 0 {
		return -1
	}

	lastIndex := len(m.contents) - 1
	extracted := m.contents[lastIndex]

	m.contents[0] = m.contents[lastIndex]
	m.contents = m.contents[:lastIndex]

	m.maxHeapifyDown()

	return extracted
}

func (m *MaxHeap) maxHeapifyUp() {
	i := len(m.contents) - 1
	for m.contents[parent(i)] < m.contents[i] {
		m.swap(parent(i), i)

		i = parent(i)
	}
}

func (m *MaxHeap) maxHeapifyDown() {

	lastIndex := len(m.contents) - 1
	parentIndex := 0
	leftChildIndex, rightChildIndex := left(parentIndex), right(parentIndex)
	childToCompareIndex := 0

	for leftChildIndex <= lastIndex {
		if leftChildIndex == lastIndex || m.contents[leftChildIndex] > m.contents[rightChildIndex] {
			childToCompareIndex = leftChildIndex
		} else {
			childToCompareIndex = rightChildIndex
		}

		if m.contents[parentIndex] < m.contents[childToCompareIndex] {
			m.swap(parentIndex, childToCompareIndex)
			parentIndex = childToCompareIndex
			leftChildIndex, rightChildIndex = left(parentIndex), right(parentIndex)
		} else {
			return
		}

	}

}

func parent(i int) int {
	return (i - 1) / 2
}

func (m *MaxHeap) swap(a, b int) {
	m.contents[a], m.contents[b] = m.contents[b], m.contents[a]
}
func left(parentIndex int) int {
	return parentIndex*2 + 1
}

func right(parentIndex int) int {
	return left(parentIndex) + 1
}
