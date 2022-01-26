//Author - Lalith Pallage
// Bubble Sort

package main

import "fmt"

func main() {
	si := []int{10, 21, 3, 56, 66, 87, 4, 78, 90, 11}
	fmt.Println("Before sorting : ", si)
	BubbleSort(si)
	fmt.Println("Sorted Slice:", si)
}

//BubbleSort function to sort a slice of integer
func BubbleSort(si []int) {
	//variables for check sorted slice and done swapping elements
	sort := false
	//get the length of slice
	length := len(si)

	//check if sort status
	for !sort {
		swapped := false

		//for loop to start the sorting
		for i := 0; i < length-1; i++ {
			//check element next is greater than the previous element
			if si[i] > si[i+1] {
				//swap element
				swapslice(&si[i+1], &si[i])
				swapped = true
			}
		}
		//check if sorted slice
		if !swapped {
			sort = true
		}
		length = length - 1
	}
}

//A function to swap elements
func swapslice(startItem, nextItem *int) {
	temp := *startItem
	*startItem = *nextItem
	*nextItem = temp
}
