package main

type Queue []interface{}

func (q *Queue) enqueue(item interface{}) {
	*q = append(*q, item)
}

func (q *Queue) dequeue() (interface{}, bool) {
	if len(*q) == 0 {
		return nil, false
	}

	item := (*q)[0]
	*q = (*q)[1:]
	return item, true
}

func (q *Queue) isEmpty() bool {
	return len(*q) == 0
}
