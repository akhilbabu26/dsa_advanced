package main

import "fmt"

func main(){
	arr := []int{8,7,5,2,4,1,3,6}
	fmt.Println("Unsorted arr: ",arr)

	BubbleSort(arr)
	fmt.Println("Bubble sort: ",arr)

	InsertionSort(arr)
	fmt.Println("Insertion sort: ",arr)

	SelectionSort(arr)
	fmt.Println("Selection sort: ",arr)

	fmt.Println("Quick sort: ",QuickSort(arr))
	
	fmt.Println("Merge sort: ",mergeSort(arr))
}

func BubbleSort(nums []int){
	l := len(nums)
	for i := 0; i<l-1; i++{
		swapped := false
		for j := 0; j < l-i-1; j++{
			if nums[j] > nums[j+1]{
				nums[j], nums[j+1] = nums[j+1], nums[j]
				swapped = true
			}
		}
		if !swapped{
			break
		}
	}
}

func InsertionSort(nums []int){
	l := len(nums)

	for i := 1; i < l; i++{
		key := nums[i]
		j := i-1

		for j >= 0 && nums[j] > key{
			nums[j+1] = nums[j]
			j--
		}

		nums[j+1] = key
	}
}

func SelectionSort(nums []int){
	l := len(nums)

	for i:=0; i<l-1; i++{
		minIndex := i

		for j:=i+1; j<l; j++{
			if nums[j] < nums[minIndex]{
				minIndex = j
			}
		}

		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}

func QuickSort(nums []int)[]int{
	if len(nums) <= 1 {
		return nums
	}

	pivot := nums[len(nums)-1]

	left := []int{}
	right := []int{}

	for i := 0; i < len(nums)-1; i++{
		if nums[i] < pivot{
			left = append(left, nums[i])
		}else{
			right = append(right, nums[i])
		}
	}

	left = QuickSort(left)
	right = QuickSort(right)

	result := append(left, pivot)
	result = append(result, right...)

	return result
}

func mergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}

	mid := len(arr) / 2

	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])

	return merge(left, right)
}

func merge(left, right []int) []int {
	result := []int{}

	i, j := 0, 0

	for i < len(left) && j < len(right) {
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}

	result = append(result, left[i:]...)
	result = append(result, right[j:]...)

	return result
}