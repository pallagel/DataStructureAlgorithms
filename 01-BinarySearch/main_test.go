//Author - Lalith Pallage
//Go unit test for unit test and benchmark

package main

import (
	"testing"
)

//define a data table to used in testing
var tests = []struct {
	data   []int
	answer int
}{
	{[]int{21, 20, 16, 7}, 20},
	{[]int{21, 20, 16, 7}, 16},
	{[]int{21, 20, 16, 7}, 19},
}

//simpel unit test running using data table for simple binary search
func TestSearchBasic(t *testing.T) {
	for _, v := range tests {
		SimpleSearch(v.data, v.answer)
	}

}

//simpel unit test running using data table for enhance Binary search
func TestEnhanceBinarySearch01(t *testing.T) {
	for _, v := range tests {
		EnhanceBinarySearch01(v.data, v.answer)
	}

}

//Benchmark test for Simple seach
func BenchmarkSimpleSearch(b *testing.B) {
	num := 50
	si := []int{1, 4, 5, 67, 87, 65, 8}
	for i := 0; i < b.N; i++ {
		SimpleSearch(si, num)
	}
}

//Benchmark test for Enhance Binary seach
func BenchmarkEnhanceBinarySearch01(b *testing.B) {
	num := 50
	si := []int{1, 4, 5, 67, 87, 65, 8}
	for i := 0; i < b.N; i++ {
		EnhanceBinarySearch01(si, num)
	}
}

//Benchmark test for Enhance Binary seach
func BenchmarkEnhanceBinarySearch02(b *testing.B) {
	num := 50
	si := []int{1, 4, 5, 67, 87, 65, 8}
	for i := 0; i < b.N; i++ {
		EnhanceBinarySearch02(si, num)
	}
}
