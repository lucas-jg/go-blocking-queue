package queue

import (
	"sync"
)

type Queue struct {
	front *node
	back  *node
	count int
	mutex *sync.Mutex
}

type node struct {
	elem     interface{}
	nextNode *node
}

func New() *Queue {

	q := &Queue{
		front: nil,
		back:  nil,
		count: 0,
		mutex: &sync.Mutex{},
	}

	return q
}

func (q *Queue) createNode(elem interface{}) *node {
	n := &node{
		elem:     elem,
		nextNode: nil,
	}

	return n
}

func (q *Queue) Enqueue(elem interface{}) {

	n := q.createNode(elem)

	if q.front == nil {
		q.front = n
		q.back = n
	} else {
		q.back.nextNode = n
		q.back = n
	}
	q.count++
}

func (q *Queue) Dequeue() interface{} {
	n := q.front
	defer q.nilPointer(n)

	if q.front.nextNode == nil {
		q.front = nil
		q.back = nil
	} else {
		q.front = q.front.nextNode
	}
	q.count--

	return n.elem
}

func (q *Queue) IsEmpty() bool {
	return q.front == nil
}

func (q *Queue) nilPointer(n *node) {
	n = nil
}
