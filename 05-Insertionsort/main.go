//Author Lalith Pallage
//Insertionsort

package main

import "fmt"

func main() {
	si := []int{10, 21, 3, 56, 66, 87, 4, 78, 90, 11}
	fmt.Println("Unsorted Slice : ", si)

	insertionsort(si)
	fmt.Println("Sorted Slice : ", si)
}

//insertionsort function
func insertionsort(si []int) {
	length := len(si)

	for i := 0; i < length; i++ {
		j := i

		for j > 0 {
			if si[j-1] > si[j] {
				swapValue(&si[j], &si[j-1])
			}
			j = j - 1
		}
	}
}

//swapValue - swap value in the slice
func swapValue(index, nextIndex *int) {
	temp := *index
	*index = *nextIndex
	*nextIndex = temp
}
