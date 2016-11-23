package main

import (
	"fmt"
	"math/rand"
)

func getMedian(arr []int) (medianIndex, median int) {

	// Intn returns, as an int, a non-negative pseudo-random number in [0,n)
	medianIndex = rand.Intn(len(arr))
	median = arr[medianIndex]
	fmt.Printf("Median: index: %d => val: %d \n", medianIndex, median)

	return
}

func QuickSort(arr []int, ch chan struct{}) {
	if ch == nil {
		return
	}
	defer func() {
		ch <- struct{}{}
	}()
	if len(arr) <= 1 {
		return
	}
	medianIndex, medianValue := getMedian(arr)

	for index, item := range arr {
		fmt.Printf("Median index: %d => index: %d \n", medianIndex, index)
		if index < medianIndex {
			switch {
			case item > medianValue:
				arr[index], arr[medianIndex] = arr[medianIndex], arr[index]
				medianIndex = index
			}
		} else {
			switch {
			case item < medianValue:
				arr[index], arr[medianIndex] = arr[medianIndex], arr[index]
				medianIndex = index
			}
		}
	}

	end := make(chan struct{})
	go QuickSort(arr[:medianIndex], end)   // Left side
	go QuickSort(arr[medianIndex+1:], end) // Right side
	<-end
	<-end
}

func main() {
	arr := []int{9, 2, 10, 12, 3, 7, 8, 4, 6, 5, 1}

	fmt.Println("Initial array is:", arr)

	end := make(chan struct{})
	go QuickSort(arr, end)
	<-end

	fmt.Println("Sorted array is:", arr)
}
