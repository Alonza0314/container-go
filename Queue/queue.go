package Queue

import (
	"errors"
)

type queue []interface{}

func NewQueue() *queue {
	q := make(queue, 0)
	return &q
}

func (q *queue) Len() int {
	return len(*q)
}

func (q *queue) IsEmpty() bool {
	return q.Len() == 0
}

func (q *queue) Front() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}
	return (*q)[0], nil
}

func (q *queue) Push(item interface{}) {
	(*q) = append((*q), item)
}

func (q *queue) Pop() (interface{}, error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty.")
	}
	front := (*q)[0]
	(*q) = (*q)[1:]
	return front, nil
}