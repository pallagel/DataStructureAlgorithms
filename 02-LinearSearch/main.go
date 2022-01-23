//Author - Lalith Pallage
//Linear Search

package main

import (
	"sort"
)

//Linear Search Function
//LinearSearch , accept slice of int and a number to search
func LinearSearch(si []int, num int) bool {
	//sort the slice
	sort.Ints(si)
	//print the sorted slice
	//fmt.Println(si)

	//for loop to loop through the slice till find the element
	for _, value := range si {
		//if value in the slice is equal to number in the search
		if value == num {
			return true
		}
	}

	return false
}
