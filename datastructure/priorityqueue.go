package datastructure

import "container/heap"

// Interface for an int priority queue. Acts as a wrapper around a general min priority
// queue and the heap interface.
type PriorityQueue interface {

	// Get the length of the queue
	Len() int

	// Pushes a value on the queue with the given priority
	Push(val interface{}, priority int)

	// Pops off the top of the queue the value with the highest priority (lowest number)
	Pop() interface{}
}

// This private struct wraps the heap.Interface implementation
type priorityQueue struct {
	valueQueue heap.Interface
}

// Creates a new PriorityQueue
func NewPriorityQueue() PriorityQueue {
	valueQ := valueQueue{}

	heap.Init(&valueQ)

	return &priorityQueue{
		valueQueue: &valueQ,
	}
}

// Returns the length of the priority queue.
func (pq priorityQueue) Len() int {
	return pq.valueQueue.Len()
}

// Adds the given value to the PQ with the given priority. Lower priority values will be
// popped off the queue first. Priority can be any integer Int.Min <= i <= Int.Max.
func (pq *priorityQueue) Push(val interface{}, priority int) {
	heap.Push(pq.valueQueue, &value{
		value:    val,
		priority: priority,
		index:    pq.valueQueue.Len(),
	})
}

// Returns the highest priority item off the PQ
func (pq *priorityQueue) Pop() interface{} {
	return heap.Pop(pq.valueQueue).(*value).value
}

// The value struct contains data required by the valueQueue and heap
type value struct {
	value    interface{}
	priority int
	index    int
}

// Internal queue type that implements the sort.Interface and heap.Interface, thus allowing it
// to be used with the heap functions.
type valueQueue []*value

// (Sort interface)
func (valQ valueQueue) Len() int {
	return len(valQ)
}

// Less reports whether the element with index i should sort before the element with index j.
// (Sort interface)
func (valQ valueQueue) Less(i, j int) bool {
	return valQ[i].priority < valQ[j].priority
}

// (Sort interface)
func (valQ valueQueue) Swap(i, j int) {
	valQ[i], valQ[j] = valQ[j], valQ[i]
}

// Adds a Value struct to the valueQueue
// (Heap interface)
func (valQ *valueQueue) Push(x interface{}) {
	valPtr := x.(*value)
	valPtr.index = len(*valQ) // set to send of queue
	*valQ = append(*valQ, valPtr)
}

// Removes the highest priority Value struct from the valueQueue
// (Heap interface)
func (valQ *valueQueue) Pop() interface{} {
	queue := *valQ
	n := len(queue)
	val := queue[n-1]
	queue[n-1] = nil       // kill reference to item
	val.index = -1         // Just in case,
	*valQ = queue[0 : n-1] // reduce size
	return val
}
