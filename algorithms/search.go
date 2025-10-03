package algorithms

import "sort"

func LinearSearch(arr []int, target int) int {
	for i, v := range arr {
		if v == target {
			return i
		}
	}
	return -1
}

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	
	for left <= right {
		mid := left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func BinarySearchSorted(arr []int, target int) int {
	return BinarySearch(arr, target)
}

func BinarySearchUnsorted(arr []int, target int) int {
	sorted := make([]int, len(arr))
	copy(sorted, arr)
	sort.Ints(sorted)
	
	index := BinarySearch(sorted, target)
	if index == -1 {
		return -1
	}
	
	for i, v := range arr {
		if v == sorted[index] {
			return i
		}
	}
	return -1
}
