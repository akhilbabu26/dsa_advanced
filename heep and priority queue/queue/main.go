package main

import "fmt"

type Item struct {
	value    string
	priority int
}

type PriorityQueue struct {
	data []Item
}

func (pq *PriorityQueue) Push(item Item) {
	pq.data = append(pq.data, item)
	pq.bubbleUp(len(pq.data) - 1)
}

func (pq *PriorityQueue) bubbleUp(i int) {
	for i > 0 {
		p := (i - 1) / 2
		if pq.data[p].priority <= pq.data[i].priority {
			break // parent has higher/equal priority, stop
		}
		pq.data[p], pq.data[i] = pq.data[i], pq.data[p]
		i = p
	}
}

func (pq *PriorityQueue) Pop() Item {
	top := pq.data[0]
	last := len(pq.data) - 1
	pq.data[0] = pq.data[last]
	pq.data = pq.data[:last]
	pq.bubbleDown(0)
	return top
}

func (pq *PriorityQueue) bubbleDown(i int) {
	n := len(pq.data)
	for {
		smallest := i
		l, r := 2*i+1, 2*i+2
		if l < n && pq.data[l].priority < pq.data[smallest].priority {
			smallest = l
		}
		if r < n && pq.data[r].priority < pq.data[smallest].priority {
			smallest = r
		}
		if smallest == i {
			break
		}
		pq.data[i], pq.data[smallest] = pq.data[smallest], pq.data[i]
		i = smallest
	}
}

func (pq *PriorityQueue) Size() int { return len(pq.data) }

func main() {
	pq := &PriorityQueue{}

	// Lower number = higher priority (like a hospital triage)
	pq.Push(Item{"low fever", 3})
	pq.Push(Item{"heart attack", 1})
	pq.Push(Item{"broken arm", 2})
	pq.Push(Item{"headache", 4})

	fmt.Println("Treating patients in priority order:")
	for pq.Size() > 0 {
		item := pq.Pop()
		fmt.Printf("  priority %d → %s\n", item.priority, item.value)
	}
}