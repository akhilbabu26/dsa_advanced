package main

import "fmt"

func main(){
	nums := []int{8,5,2,3,6,9,7,4,1}

	BubbleSort(nums)
	fmt.Println(nums)

	InsertionSort(nums)
	fmt.Println(nums)
	
	SelectionSort(nums)
	fmt.Println(nums)
}

func BubbleSort(nums []int){
	n := len(nums)

	for i := 0; i < n-1; i++{
		swapped := false
		for j := 0; j < n-i-1; j++{
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
	n := len(nums)
	
	for i:=1; i<n; i++{
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
	n := len(nums)

	for i:=0; i < n-1; i++{
		minIndex := i

		for j:=i+1; j<n; j++{
			if nums[j] < nums[minIndex]{
				minIndex = j
			}
		}
		nums[i], nums[minIndex] = nums[minIndex], nums[i]
	}
}