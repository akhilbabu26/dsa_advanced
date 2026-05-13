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

	// bubble 
	for i > 0 {

		parent := (i - 1) / 2

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

	h.data[0] = h.data[last]

	h.data = h.data[:last]

	i := 0
	n := len(h.data)

	// bubble down
	for {

		smallest := i

		left := 2*i + 1
		right := 2*i + 2

		if left < n && h.data[left] < h.data[smallest] {

			smallest = left
		}

		if right < n && h.data[right] < h.data[smallest] {

			smallest = right
		}

		if smallest == i {
			break
		}

		h.data[i], h.data[smallest] = h.data[smallest], h.data[i]

		i = smallest
	}

	return min, true
}

// HEAPIFY
func (h *MinHeap) Heapify(arr []int) {

	h.data = make([]int, len(arr))
	copy(h.data, arr)

	n := len(h.data)

	for i := n/2 - 1; i >= 0; i-- {

		current := i

		for {

			smallest := current

			left := 2*current + 1
			right := 2*current + 2

			if left < n && h.data[left] < h.data[smallest] {
				smallest = left
			}

			if right < n && h.data[right] < h.data[smallest] {
				smallest = right
			}

			if smallest == current {
				break
			}

			h.data[current], h.data[smallest] = h.data[smallest], h.data[current]
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
		if h.data[parent] >= h.data[i] {
			break
		}
		// swap
		h.data[parent], h.data[i] = h.data[i], h.data[parent]

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
	h.data[0] = h.data[last]
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

		h.data[i], h.data[largest] = h.data[largest], h.data[i]
		i = largest
	}
	return max, true
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

			h.data[current], h.data[largest] = h.data[largest], h.data[current]
			current = largest
		}
	}
}

// ================= MAIN =================
func main() {
	// ================= MIN HEAP =================
	minH := &MinHeap{}

	arr := []int{9, 3, 7, 1, 5, 8, 2}

	for _, v := range arr {
		minH.Insert(v)
	}
	fmt.Println(minH.data)

	// ================= MAX HEAP =================
	maxH := &MaxHeap{}
	arr1 := []int{9, 3, 7, 1, 5, 8, 2}

	for _, v := range arr1 {
		maxH.Insert(v)
	}
	fmt.Println(maxH.data)
}