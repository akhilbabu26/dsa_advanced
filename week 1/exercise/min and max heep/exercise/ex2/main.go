package main
import "fmt"

type MinHeap struct {
	data []int
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	i := len(h.data) - 1

	for i > 0 {
		parent := (i - 1) / 2

		if h.data[parent] <= h.data[i] {
			break
		}

		h.data[parent], h.data[i] = h.data[i], h.data[parent]

		i = parent
	}
}

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

func (h *MinHeap) ExtractMin() int {
	if len(h.data) == 0 {
		return -1
	}

	min := h.data[0]
	lastIndex := len(h.data) - 1
	h.data[0] = h.data[lastIndex]
	h.data = h.data[:lastIndex]
	n := len(h.data)

	current := 0

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

	return min
}

func main() {
	arr := []int{4, 7, 5, 2, 1, 6, 8, 9, 9, 3, 0}

	minH := &MinHeap{}
	for _, v := range arr {
		minH.Insert(v)
	}
	fmt.Println(minH.data)

	// Heapify
	minH.Heapify(arr)
	fmt.Println(minH.data)

	// Extract Min
    minH.ExtractMin()
	fmt.Println(minH.data)
}