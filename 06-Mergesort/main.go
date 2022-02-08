//Author - Lalith Pallage
//Merge sort - First attempt

package main

import (
	"fmt"
)

func main() {
	si := []int{1, 4, 3, 2, 10, 0, 12, 21}
	fmt.Println("Before Sort :", si)
	newSI := divide(si)
	fmt.Println("After Sort :", newSI)
}

//divide Split the slice in to equal (almost) parts
func divide(si []int) []int {
	//return slice if length is 1
	if len(si) == 1 {
		return si
	}

	//split lef and right side
	left := divide(si[:len(si)/2])
	right := divide(si[len(si)/2:])
	fmt.Println("LEFT ", left)
	fmt.Println("RIGHT ", right)

	//call to the mergeSort funcion
	return mergeSort(left, right)
}

//mergeSort - function to sort the split slice/array
func mergeSort(left, right []int) []int {
	result := make([]int, 0)
	iPtr, jPtr := 0, 0

	for iPtr < len(left) && jPtr < len(right) {
		if left[iPtr] < right[jPtr] {
			result = append(result, left[iPtr])
			fmt.Println("IF", result)
			iPtr++
		} else {
			result = append(result, right[jPtr])
			fmt.Println("ELSE", result)
			jPtr++
		}
	}

	for ; iPtr < len(left); iPtr++ {
		result = append(result, left[iPtr])
		fmt.Println("First FOR", result)

	}
	for ; jPtr < len(right); jPtr++ {
		result = append(result, right[jPtr])
		fmt.Println("Second FOR ", result)
	}

	return result
}
