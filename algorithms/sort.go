package algorithms

import (
	"math/rand"
	"sort"
	"time"
)

func BubbleSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	
	n := len(result)
	for i := 0; i < n-1; i++ {
		swapped := false
		for j := 0; j < n-i-1; j++ {
			if result[j] > result[j+1] {
				result[j], result[j+1] = result[j+1], result[j]
				swapped = true
			}
		}
		if !swapped {
			break
		}
	}
	return result
}

func InsertionSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	
	for i := 1; i < len(result); i++ {
		key := result[i]
		j := i - 1
		for j >= 0 && result[j] > key {
			result[j+1] = result[j]
			j--
		}
		result[j+1] = key
	}
	return result
}

func MergeSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	mid := len(arr) / 2
	left := MergeSort(arr[:mid])
	right := MergeSort(arr[mid:])
	
	return merge(left, right)
}

func merge(left, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	i, j := 0, 0
	
	for i < len(left) && j < len(right) {
		if left[i] <= right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	
	result = append(result, left[i:]...)
	result = append(result, right[j:]...)
	
	return result
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	
	result := make([]int, len(arr))
	copy(result, arr)
	
	rand.Seed(time.Now().UnixNano())
	quickSortHelper(result, 0, len(result)-1)
	
	return result
}

func quickSortHelper(arr []int, low, high int) {
	if low < high {
		pi := partition(arr, low, high)
		quickSortHelper(arr, low, pi-1)
		quickSortHelper(arr, pi+1, high)
	}
}

func partition(arr []int, low, high int) int {
	pivot := arr[high]
	i := low - 1
	
	for j := low; j < high; j++ {
		if arr[j] <= pivot {
			i++
			arr[i], arr[j] = arr[j], arr[i]
		}
	}
	
	arr[i+1], arr[high] = arr[high], arr[i+1]
	return i + 1
}

func HeapSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	
	n := len(result)
	
	for i := n/2 - 1; i >= 0; i-- {
		heapify(result, n, i)
	}
	
	for i := n - 1; i > 0; i-- {
		result[0], result[i] = result[i], result[0]
		heapify(result, i, 0)
	}
	
	return result
}

func heapify(arr []int, n, i int) {
	largest := i
	left := 2*i + 1
	right := 2*i + 2
	
	if left < n && arr[left] > arr[largest] {
		largest = left
	}
	
	if right < n && arr[right] > arr[largest] {
		largest = right
	}
	
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapify(arr, n, largest)
	}
}

func NativeSort(arr []int) []int {
	result := make([]int, len(arr))
	copy(result, arr)
	sort.Ints(result)
	return result
}
