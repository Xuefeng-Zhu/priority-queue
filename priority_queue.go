package priority_queue

type PriorityQueue struct {
	priorities  []uint64
	priorityMap map[uint64][]interface{}
}

// create a new priority queue
func NewPriorityQueue() *PriorityQueue {
	return &PriorityQueue{
		priorityMap: make(map[uint64][]interface{}),
	}
}

// check whether the queue has no elements
func (pq *PriorityQueue) IsEmpty() bool {
	return len(pq.priorities) == 0
}

// returns the highest-priority element but does not modify the queue
func (pq *PriorityQueue) Peek() interface{} {
	if pq.IsEmpty() {
		return nil
	}

	return pq.priorityMap[pq.priorities[0]][0]
}

// remove the element from the queue that has the highest priority, and return it
func (pq *PriorityQueue) Pop() (res interface{}) {
	res = pq.Peek()
	if res == nil {
		return
	}

	priority := pq.priorities[0]
	pq.priorityMap[priority] = pq.priorityMap[priority][1:]
	if len(pq.priorityMap[priority]) == 0 {
		pq.priorities = pq.priorities[1:]
	}

	return
}

// add an element to the queue with an associated priority
func (pq *PriorityQueue) Insert(element interface{}, priority uint64) {
	pq.priorityMap[priority] = append(pq.priorityMap[priority], element)

	// properly intert priority into priorities list
	for i, p := range pq.priorities {
		if priority == p {
			return
		}

		if priority > p {
			pq.priorities = append(pq.priorities[:i], append([]uint64{priority}, pq.priorities[i:]...)...)
			return
		}
	}

	pq.priorities = append(pq.priorities, priority)
}
