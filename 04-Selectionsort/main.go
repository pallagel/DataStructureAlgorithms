//Author - Lalith Pallage
//Selection sort

package main

import "fmt"

func main() {
	si := []int{10, 21, 3, 56, 66, 87, 4, 78, 90, 11}
	fmt.Println("Unsorted Slice : ", si)

	SelectionsortSecond(si)
	fmt.Println("Sorted Slice : ", si)

}

//Selectionsort - Function to sort the slice
func Selectionsort(si []int) {
	length := len(si)

	//variable to store current min value and variable for 'for' loop
	var currentMin, i, j int

	//set first index as current minumun
	for j = 0; j < length-1; j++ {
		currentMin = j

		//loop through slice to find next min value
		for i = j + 1; i < length; i++ {
			if si[i] < si[currentMin] {
				currentMin = i
			}
		}

		//swap values with the min value
		if currentMin != j {
			swapValue(&si[j], &si[currentMin])
		}
	}
}

//swapValue - swap value in the slice
func swapValue(index, nextIndex *int) {
	temp := *index
	*index = *nextIndex
	*nextIndex = temp
}

//Second method to do the same
//SelectionsortSecond
func SelectionsortSecond(si []int) {
	length := len(si)

	//set first index as current minumun
	for i := 0; i < length-1; i++ {
		currentMin := i

		//loop through slice to find next min value
		for j := i + 1; j < length; j++ {
			if si[j] < si[currentMin] {
				currentMin = j
			}
		}

		//swap values with the min value
		if currentMin != i {
			si[i], si[currentMin] = si[currentMin], si[i]
		}
	}
}
