package main

import "fmt"

// ─────────────────────────────────────────
// MIN HEAP
// ─────────────────────────────────────────

type MinHeap struct {
	data []int
}

// index helpers
func (h *MinHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MinHeap) left(i int) int   { return 2*i + 1 }
func (h *MinHeap) right(i int) int  { return 2*i + 2 }
func (h *MinHeap) size() int        { return len(h.data) }

// Insert — add to end, bubble UP — O(log n)
func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

func (h *MinHeap) bubbleUp(i int) {
	for i > 0 && h.data[h.parent(i)] > h.data[i] {
		p := h.parent(i)
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

// ExtractMin — remove root, put last at top, bubble DOWN — O(log n)
func (h *MinHeap) ExtractMin() (int, bool) {
	if h.size() == 0 {
		return 0, false
	}
	min := h.data[0]
	last := h.size() - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if h.size() > 0 {
		h.bubbleDown(0)
	}
	return min, true
}

func (h *MinHeap) bubbleDown(i int) {
	n := h.size()
	for {
		smallest := i
		l, r := h.left(i), h.right(i)
		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if smallest == i {
			break
		}
		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]
		i = smallest
	}
}

// Heapify — build heap from existing slice — O(n)
// Floyd's algorithm: bubble down from last non-leaf to root
func (h *MinHeap) Heapify(data []int) {
	h.data = make([]int, len(data))
	copy(h.data, data)
	// last non-leaf index = (n/2) - 1
	for i := h.size()/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

func (h *MinHeap) Peek() (int, bool) {
	if h.size() == 0 {
		return 0, false
	}
	return h.data[0], true
}

// ─────────────────────────────────────────
// MAX HEAP
// ─────────────────────────────────────────

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) parent(i int) int { return (i - 1) / 2 }
func (h *MaxHeap) left(i int) int   { return 2*i + 1 }
func (h *MaxHeap) right(i int) int  { return 2*i + 2 }
func (h *MaxHeap) size() int        { return len(h.data) }

// Insert — add to end, bubble UP — O(log n)
func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.bubbleUp(len(h.data) - 1)
}

func (h *MaxHeap) bubbleUp(i int) {
	for i > 0 && h.data[h.parent(i)] < h.data[i] {
		// parent is SMALLER than child — swap (opposite of minheap)
		p := h.parent(i)
		h.data[p], h.data[i] = h.data[i], h.data[p]
		i = p
	}
}

// ExtractMax — remove root, put last at top, bubble DOWN — O(log n)
func (h *MaxHeap) ExtractMax() (int, bool) {
	if h.size() == 0 {
		return 0, false
	}
	max := h.data[0]
	last := h.size() - 1
	h.data[0] = h.data[last]
	h.data = h.data[:last]
	if h.size() > 0 {
		h.bubbleDown(0)
	}
	return max, true
}

func (h *MaxHeap) bubbleDown(i int) {
	n := h.size()
	for {
		largest := i
		l, r := h.left(i), h.right(i)
		if l < n && h.data[l] > h.data[largest] {
			largest = l
		}
		if r < n && h.data[r] > h.data[largest] {
			largest = r
		}
		if largest == i {
			break
		}
		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
}

// Heapify — build heap from existing slice — O(n)
func (h *MaxHeap) Heapify(data []int) {
	h.data = make([]int, len(data))
	copy(h.data, data)
	for i := h.size()/2 - 1; i >= 0; i-- {
		h.bubbleDown(i)
	}
}

func (h *MaxHeap) Peek() (int, bool) {
	if h.size() == 0 {
		return 0, false
	}
	return h.data[0], true
}

// ─────────────────────────────────────────
// MAIN
// ─────────────────────────────────────────

func main() {
	fmt.Println("═══════════════════════════")
	fmt.Println("         MIN HEAP          ")
	fmt.Println("═══════════════════════════")

	minH := &MinHeap{}

	// Insert
	for _, v := range []int{9, 3, 7, 1, 5, 8, 2} {
		minH.Insert(v)
		fmt.Printf("Insert(%d) → heap: %v\n", v, minH.data)
	}

	// Peek
	if val, ok := minH.Peek(); ok {
		fmt.Printf("\nPeek()     → %d (min, no removal)\n", val)
	}

	// Heapify
	raw := []int{15, 4, 11, 2, 9, 6}
	minH.Heapify(raw)
	fmt.Printf("\nHeapify(%v)\n", raw)
	fmt.Printf("Result     → %v\n", minH.data)

	// ExtractMin
	fmt.Println("\nExtracting all (should be sorted ascending):")
	fmt.Print("→ ")
	for minH.size() > 0 {
		val, _ := minH.ExtractMin()
		fmt.Printf("%d ", val) // 2 4 6 9 11 15
	}
	fmt.Println()

	fmt.Println("\n═══════════════════════════")
	fmt.Println("         MAX HEAP          ")
	fmt.Println("═══════════════════════════")

	maxH := &MaxHeap{}

	// Insert
	for _, v := range []int{9, 3, 7, 1, 5, 8, 2} {
		maxH.Insert(v)
		fmt.Printf("Insert(%d) → heap: %v\n", v, maxH.data)
	}

	// Peek
	if val, ok := maxH.Peek(); ok {
		fmt.Printf("\nPeek()     → %d (max, no removal)\n", val)
	}

	// Heapify
	maxH.Heapify(raw)
	fmt.Printf("\nHeapify(%v)\n", raw)
	fmt.Printf("Result%v\n", maxH.data)

	// ExtractMax
	fmt.Println("\nExtracting all (should be sorted descending):")
	fmt.Print("→ ")
	for maxH.size() > 0 {
		val, _ := maxH.ExtractMax()
		fmt.Printf("%d ", val) // 15 11 9 6 4 2
	}
	fmt.Println()
}