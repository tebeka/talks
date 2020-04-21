package main

import (
	"container/heap"
	"fmt"
)

// Task to perform
type Task struct {
	Priority int
	Name     string
}

// PriorityQueue is a priority queue of tasks
type PriorityQueue []Task

/*
type Interface interface {
    sort.Interface
    Push(x interface{}) // add x as element Len()
    Pop() interface{}   // remove and return element Len() - 1.
}
*/

func (pq PriorityQueue) Len() int { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}
func (pq PriorityQueue) Swap(i, j int)       { pq[i], pq[j] = pq[j], pq[i] }
func (pq *PriorityQueue) Push(x interface{}) { *pq = append(*pq, x.(Task)) }
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	task := (*pq)[n-1]
	*pq = (*pq)[:n-1] // TODO: cap(*pq) > 2 * len(*pq)
	return task
}

func main() {
	var pq PriorityQueue
	heap.Push(&pq, Task{40, "Teach Go"})
	heap.Push(&pq, Task{20, "Write Book"})
	heap.Push(&pq, Task{50, "Update OS"})
	heap.Push(&pq, Task{10, "Wake Up"})
	heap.Push(&pq, Task{30, "Feed Cats"})

	for len(pq) > 0 {
		task := heap.Pop(&pq)
		fmt.Printf("%+v\n", task)
	}
}
