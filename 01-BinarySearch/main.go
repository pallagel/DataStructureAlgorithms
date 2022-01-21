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
	//sory the slice
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

func main() {
	//declare a slice
	numbers := []int{23, 4, 5, 67, 8, 99, 27}
	//call the function to find the value
	fmt.Println("Found the number \t\t:", SimpleBinarySearch(numbers, 67))             //print true
	fmt.Println("The number is not in the slice \t:", SimpleBinarySearch(numbers, 69)) //print false
}
