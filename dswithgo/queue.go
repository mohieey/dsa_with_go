package dswithgo

import (
	"errors"
	"fmt"
)

type Queue struct {
	data   []int
	Length int
}

func (q *Queue) Enqueue(num int) {
	q.data = append(q.data, num)
	q.Length = len(q.data)
}

func (q *Queue) Dequeue() (int, error) {
	if q.Length == 0 {
		return 0, errors.New("the queue is empty")
	}
	dequeuedValue := q.data[0]
	q.data = q.data[1:]
	q.Length = len(q.data)
	return dequeuedValue, nil
}

func (q *Queue) Print() {
	fmt.Println(q.data, " Length: ", q.Length)
}
