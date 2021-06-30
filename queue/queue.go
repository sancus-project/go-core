// thread safe queue without slices
package queue

import (
	"sync"
)

type entry struct {
	v    interface{}
	next *entry
}

type Queue struct {
	mu   sync.Mutex
	next *entry
	last *entry
}

func (q *Queue) Push(v interface{}) {
	p := &entry{v: v}

	q.mu.Lock()
	defer q.mu.Unlock()

	if q.last == nil {
		q.next = p
		q.last = p
	} else {
		q.last.next = p
		q.last = p
	}
}

func (q *Queue) Pop() (interface{}, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if p := q.next; p != nil {
		q.next = p.next
		return p.v, true
	}

	return nil, false
}
