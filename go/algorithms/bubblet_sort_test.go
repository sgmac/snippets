package main

import "testing"

func BenchmarkBubbleA(b *testing.B) {
	arrays := [][]int{
		{1, 2, 3, 4, 5},
		{3, 2, 3, 4, 5},
		{5, 3, 2, 1, 3},
	}

	for _, a := range arrays {
		for n := 0; n < b.N; n++ {
			Bubble(a)
		}
	}
}

func BenchmarkBubbleB(b *testing.B) {
	arrays := [][]int{
		{1, 2, 3, 4, 5},
		{3, 2, 3, 4, 5},
		{5, 3, 2, 1, 3},
	}

	for _, a := range arrays {
		for n := 0; n < b.N; n++ {
			BubbleSorted(a)
		}
	}
}
