package priority_queue

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestEmpty(t *testing.T) {
	priorityQueue := NewPriorityQueue()

	assert.Equal(t, true, priorityQueue.IsEmpty())
	assert.Equal(t, nil, priorityQueue.Peek())
	assert.Equal(t, nil, priorityQueue.Pop())
}

func TestInsert(t *testing.T) {
	priorityQueue := NewPriorityQueue()
	priorityQueue.Insert(1, 1)

	assert.Equal(t, false, priorityQueue.IsEmpty())
	assert.Equal(t, 1, priorityQueue.Peek())
	assert.Equal(t, 1, priorityQueue.Pop())
	assert.Equal(t, true, priorityQueue.IsEmpty())
}

func TestBasic(t *testing.T) {
	priorityQueue := NewPriorityQueue()
	priorityQueue.Insert(2, 2)
	priorityQueue.Insert(1, 1)
	priorityQueue.Insert(3, 3)
	priorityQueue.Insert(2, 2)

	assert.Equal(t, 3, priorityQueue.Pop())
	assert.Equal(t, 2, priorityQueue.Pop())
	assert.Equal(t, 2, priorityQueue.Pop())
	assert.Equal(t, 1, priorityQueue.Pop())
	assert.Equal(t, true, priorityQueue.IsEmpty())
}
