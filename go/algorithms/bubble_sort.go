package main

import "fmt"

/*
 non-increasing vs increasing:  the last one implies the next number is
 always greater than the previous one.

 Bubble sort compares each consecutive pair of numbers.
 If the number is greater than the second one, swap position.

 5 4 2 3 1 0 -> 4 5 2 3 1 0
 4 5 2 3 1 0 -> 4 2 5 3 1 0
 4 2 5 3 1 0 -> 4 2 3 5 1 0
 4 2 3 5 1 0 -> 4 2 3 1 5 0
 4 2 3 1 5 0 -> 4 2 3 1 0 5

 This is the first walk throught the list.
 N * N-1  iterations, don't compare the last number.
*/

func main() {
	unorderedList := []int{5, 4, 2, 3, 1, 0}
	orderedList := []int{0, 1, 2, 3, 4, 5}
	fmt.Printf("unorderedList before bubble: %v\n", unorderedList)
	Bubble(unorderedList)
	fmt.Printf("unorderedList after bubble: %v\n", unorderedList)

	fmt.Printf("orderedList before bubble: %v\n", orderedList)
	BubbleSorted(orderedList)
	fmt.Printf("orderedList after bubble: %v\n", orderedList)
}

// First implmentation, does compare all the elements
// even the last ones after some iterations are already
// in place.
func Bubble(list []int) {
	for k := 0; k < len(list); k++ {
		for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func BubbleSorted(list []int) {
	// Verify the list is unordered
	var ordered int
	for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
		if list[i] < list[j] {
			ordered++
			continue
		}
	}
	if ordered == len(list)-1 {
		return
	}
	for k := 0; k < len(list); k++ {
		for i, j := 0, 1; j < len(list)-i; i, j = i+1, j+1 {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}

func BubbleReverse(list []int) {
	// Verify the list is unordered
	var ordered int
	for i, j := 0, 1; j < len(list); i, j = i+1, j+1 {
		if list[i] < list[j] {
			ordered++
			continue
		}
	}
	if ordered == len(list)-1 {
		return
	}
	for k := 0; k < len(list); k++ {
		for i, j := 0, 1; j < len(list)-i; i, j = i+1, j+1 {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
