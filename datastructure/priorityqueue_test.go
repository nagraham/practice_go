package datastructure

import (
	"container/heap"
	"github.com/stretchr/testify/assert"
	"testing"
)

// Test the Priority Queue

func TestPriorityQueue_Len(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Push("a", 1)
	pq.Push("b", 2)
	pq.Push("c", 3)

	assert.Equal(t, 3, pq.Len())
}

func TestPriorityQueue_PushAndPop(t *testing.T) {
	pq := NewPriorityQueue()
	pq.Push("a", 3)
	pq.Push("b", 1)
	pq.Push("c", 2)

	assert.Equal(t, "b", pq.Pop())
	assert.Equal(t, "c", pq.Pop())
	assert.Equal(t, "a", pq.Pop())
}

// Test the internal value queue

func TestValueQueueBasics(t *testing.T) {
	queue := valueQueue{
		{value: "a", priority: 1, index: 0},
		{value: "b", priority: 2, index: 1},
		{value: "c", priority: 0, index: 2},
	}

	heap.Init(&queue)

	assert.Equal(t, "c", heap.Pop(&queue).(*value).value)
	assert.Equal(t, "a", heap.Pop(&queue).(*value).value)
	assert.Equal(t, "b", heap.Pop(&queue).(*value).value)
}

func TestValueQueuePushNew(t *testing.T) {
	queue := valueQueue{
		{value: "a", priority: 5, index: 0},
	}

	heap.Init(&queue)

	heap.Push(&queue, &value{value: "b", priority: 4, index: 1})

	assert.Equal(t, "b", heap.Pop(&queue).(*value).value)
	assert.Equal(t, "a", heap.Pop(&queue).(*value).value)
}

func TestValueQueuePushPopPush(t *testing.T) {
	queue := valueQueue{
		{value: "a", priority: 5, index: 0},
	}
	heap.Init(&queue)

	heap.Push(&queue, &value{value: "b", priority: 4, index: len(queue)})
	assert.Equal(t, "b", heap.Pop(&queue).(*value).value)

	heap.Push(&queue, &value{value: "c", priority: 7, index: len(queue)})
	heap.Push(&queue, &value{value: "d", priority: 3, index: len(queue)})
	assert.Equal(t, "d", heap.Pop(&queue).(*value).value)
	assert.Equal(t, "a", heap.Pop(&queue).(*value).value)
	assert.Equal(t, "c", heap.Pop(&queue).(*value).value)
}
