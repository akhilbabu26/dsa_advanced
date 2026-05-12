package main

import "fmt"

// ================= MIN HEAP =================
type MinHeap struct {
	data []int
}

// INSERT
func (h *MinHeap) Insert(val int) {

	h.data = append(h.data, val)

	i := len(h.data) - 1

	// bubble up
	for i > 0 {

		parent := (i - 1) / 2

		// parent already smaller
		if h.data[parent] <= h.data[i] {
			break
		}

		// swap
		h.data[parent], h.data[i] =
			h.data[i], h.data[parent]

		i = parent
	}
}

// EXTRACT MIN
func (h *MinHeap) ExtractMin() (int, bool) {

	if len(h.data) == 0 {
		return 0, false
	}

	min := h.data[0]

	last := len(h.data) - 1

	// move last element to root
	h.data[0] = h.data[last]

	// remove last
	h.data = h.data[:last]

	i := 0
	n := len(h.data)

	// bubble down
	for {

		smallest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < n &&
			h.data[left] < h.data[smallest] {

			smallest = left
		}

		if right < n &&
			h.data[right] < h.data[smallest] {

			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] =
			h.data[smallest], h.data[i]
			
		i = smallest
	}

	return min, true
}

// PEEK
func (h *MinHeap) Peek() (int, bool) {

	if len(h.data) == 0 {
		return 0, false
	}

	return h.data[0], true
}

// HEAPIFY
func (h *MinHeap) Heapify(arr []int) {

	h.data = make([]int, len(arr))
	copy(h.data, arr)

	n := len(h.data)

	// last non-leaf node
	for i := n/2 - 1; i >= 0; i-- {

		current := i

		for {

			smallest := current

			left := 2*current + 1
			right := 2*current + 2

			if left < n &&
				h.data[left] < h.data[smallest] {

				smallest = left
			}

			if right < n &&
				h.data[right] < h.data[smallest] {

				smallest = right
			}

			if smallest == current {
				break
			}

			h.data[current], h.data[smallest] =
				h.data[smallest], h.data[current]

			current = smallest
		}
	}
}

// ================= MAX HEAP =================

type MaxHeap struct {
	data []int
}

// INSERT
func (h *MaxHeap) Insert(val int) {

	h.data = append(h.data, val)

	i := len(h.data) - 1

	// bubble up
	for i > 0 {

		parent := (i - 1) / 2

		// parent already larger
		if h.data[parent] >= h.data[i] {
			break
		}

		// swap
		h.data[parent], h.data[i] =
			h.data[i], h.data[parent]

		i = parent
	}
}

// EXTRACT MAX
func (h *MaxHeap) ExtractMax() (int, bool) {

	if len(h.data) == 0 {
		return 0, false
	}

	max := h.data[0]

	last := len(h.data) - 1

	// move last element to root
	h.data[0] = h.data[last]

	// remove last
	h.data = h.data[:last]

	i := 0
	n := len(h.data)

	// bubble down
	for {

		largest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < n &&
			h.data[left] > h.data[largest] {

			largest = left
		}

		if right < n &&
			h.data[right] > h.data[largest] {

			largest = right
		}

		if largest == i {
			break
		}

		h.data[i], h.data[largest] =
			h.data[largest], h.data[i]

		i = largest
	}

	return max, true
}

// PEEK
func (h *MaxHeap) Peek() (int, bool) {

	if len(h.data) == 0 {
		return 0, false
	}

	return h.data[0], true
}

// HEAPIFY
func (h *MaxHeap) Heapify(arr []int) {

	h.data = make([]int, len(arr))
	copy(h.data, arr)

	n := len(h.data)

	for i := n/2 - 1; i >= 0; i-- {

		current := i

		for {

			largest := current

			left := 2*current + 1
			right := 2*current + 2

			if left < n &&
				h.data[left] > h.data[largest] {

				largest = left
			}

			if right < n &&
				h.data[right] > h.data[largest] {

				largest = right
			}

			if largest == current {
				break
			}

			h.data[current], h.data[largest] =
				h.data[largest], h.data[current]

			current = largest
		}
	}
}

// ================= MAIN =================
func main() {

	// ================= MIN HEAP =================

	fmt.Println("===== MIN HEAP =====")

	minH := &MinHeap{}

	for _, v := range []int{9, 3, 7, 1, 5, 8, 2} {

		minH.Insert(v)

		fmt.Printf(
			"Insert(%d) -> %v\n",
			v,
			minH.data,
		)
	}

	if val, ok := minH.Peek(); ok {
		fmt.Println("\nPeek:", val)
	}

	fmt.Println("\nExtracting Min:")

	for len(minH.data) > 0 {

		val, _ := minH.ExtractMin()

		fmt.Printf(
			"ExtractMin() -> %-2d Heap: %v\n",
			val,
			minH.data,
		)
	}

	// ================= MAX HEAP =================

	fmt.Println("\n===== MAX HEAP =====")

	maxH := &MaxHeap{}

	for _, v := range []int{9, 3, 7, 1, 5, 8, 2} {

		maxH.Insert(v)

		fmt.Printf(
			"Insert(%d) -> %v\n",
			v,
			maxH.data,
		)
	}

	if val, ok := maxH.Peek(); ok {
		fmt.Println("\nPeek:", val)
	}

	fmt.Println("\nExtracting Max:")

	for len(maxH.data) > 0 {

		val, _ := maxH.ExtractMax()

		fmt.Printf(
			"ExtractMax() -> %-2d Heap: %v\n",
			val,
			maxH.data,
		)
	}
}