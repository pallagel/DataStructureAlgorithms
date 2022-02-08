package main

import "fmt"

func main() {
	fmt.Println("Hello")
	si := []int{1, 2, 4, 5, 6, 6, 7, 7}

	start := 0
	end := len(si) - 1
	middle := end / 2

	fmt.Println(si[start:middle])
	fmt.Println(si[middle:end])

	for _, i := range si[middle:end] {
		fmt.Println(si[middle : i-1])
	}
}
