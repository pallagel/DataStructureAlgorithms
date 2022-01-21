//Author - Lalith Pallage
//Binary Search in 3 diffent ways

package main

import (
	"fmt"
	"sort"
)

//Simple Function for Binary search
//SimpleBinarySearch , accept slice of int and a number to search
func SimpleBinarySearch(si []int, num int) bool {
	//sort the slice
	sort.Ints(si)
	//print the sorted slice
	fmt.Println(si)

	//for loop to loop through the slice till find the element
	for _, value := range si {
		//if value in the slice is equal to number in the search
		if value == num {
			return true
		}
	}

	return false
}

//Enhance Binary Seach function 01
//EnhanceBinarySearch using slice of slices
func EnhanceBinarySearch(si []int, num int) bool {
	//sort the int array
	sort.Ints(si)

	//start index
	startIndex := 0
	//end index of the slie
	endIndex := len(si)
	//find the middle of the slice
	middleIndex := (startIndex + endIndex) / 2

	//check the value in the middle on the slice
	if si[middleIndex] > num {
		//print the slice start to middle index
		fmt.Println(si[:middleIndex])
		//range ove the slice of slice
		for _, value := range si[:middleIndex] {
			//compare the number to vlaue in the slice
			if value == num {
				return true
			}
		}
	} else {
		//print the middle index to end of slice
		fmt.Println(si[middleIndex:])
		//range ove the slice of slice
		for _, value := range si[middleIndex:] {
			//compare the number to vlaue in the slice
			if value == num {
				return true
			}
		}
	}
	return false
}

func main() {
	//declare a slice
	numbers := []int{23, 4, 5, 67, 8, 99, 27}
	//call the function to find the value
	fmt.Println("Found the number \t\t:", SimpleBinarySearch(numbers, 67))             //print true
	fmt.Println("The number is not in the slice \t:", SimpleBinarySearch(numbers, 69)) //print false

	fmt.Println("*********** Enhance Binary Search 1 *******************")
	fmt.Println("Found the number \t\t:", EnhanceBinarySearch(numbers, 5))               //print true
	fmt.Println("The number is not in the slice \t:", EnhanceBinarySearch(numbers, 100)) //print false
}
