package algorithms

import (
	"reflect"
	"sort"
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	
	tests := []struct {
		target int
		expected int
	}{
		{5, 2},
		{1, 0},
		{15, 7},
		{4, -1},
		{20, -1},
	}
	
	for _, test := range tests {
		result := LinearSearch(arr, test.target)
		if result != test.expected {
			t.Errorf("LinearSearch(%v, %d) = %d, expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestBinarySearch(t *testing.T) {
	arr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	
	tests := []struct {
		target int
		expected int
	}{
		{5, 2},
		{1, 0},
		{15, 7},
		{4, -1},
		{20, -1},
	}
	
	for _, test := range tests {
		result := BinarySearch(arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearch(%v, %d) = %d, expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestBinarySearchUnsorted(t *testing.T) {
	arr := []int{5, 1, 9, 3, 7, 11, 15, 13}
	
	tests := []struct {
		target int
		expected int
	}{
		{5, 0},
		{1, 1},
		{15, 6},
		{4, -1},
		{20, -1},
	}
	
	for _, test := range tests {
		result := BinarySearchUnsorted(arr, test.target)
		if result != test.expected {
			t.Errorf("BinarySearchUnsorted(%v, %d) = %d, expected %d", arr, test.target, result, test.expected)
		}
	}
}

func TestBubbleSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := BubbleSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("BubbleSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestInsertionSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := InsertionSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("InsertionSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestMergeSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := MergeSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("MergeSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestQuickSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := QuickSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("QuickSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestHeapSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := HeapSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("HeapSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestNativeSort(t *testing.T) {
	arr := []int{64, 34, 25, 12, 22, 11, 90}
	expected := []int{11, 12, 22, 25, 34, 64, 90}
	
	result := NativeSort(arr)
	
	if !reflect.DeepEqual(result, expected) {
		t.Errorf("NativeSort(%v) = %v, expected %v", arr, result, expected)
	}
}

func TestSortingAlgorithmsConsistency(t *testing.T) {
	testCases := [][]int{
		{1},
		{2, 1},
		{3, 1, 2},
		{5, 2, 8, 1, 9, 3, 7, 4, 6},
		{9, 8, 7, 6, 5, 4, 3, 2, 1},
		{1, 2, 3, 4, 5, 6, 7, 8, 9},
	}
	
	algorithms := []func([]int) []int{
		BubbleSort,
		InsertionSort,
		MergeSort,
		QuickSort,
		HeapSort,
		NativeSort,
	}
	
	for _, testCase := range testCases {
		expected := make([]int, len(testCase))
		copy(expected, testCase)
		sort.Ints(expected)
		
		for _, algorithm := range algorithms {
			result := algorithm(testCase)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Algorithm %T failed for input %v: got %v, expected %v", algorithm, testCase, result, expected)
			}
		}
	}
}

func TestEmptyAndSingleElement(t *testing.T) {
	algorithms := []func([]int) []int{
		BubbleSort,
		InsertionSort,
		MergeSort,
		QuickSort,
		HeapSort,
		NativeSort,
	}
	
	testCases := [][]int{
		{},
		{42},
	}
	
	for _, testCase := range testCases {
		expected := make([]int, len(testCase))
		copy(expected, testCase)
		sort.Ints(expected)
		
		for _, algorithm := range algorithms {
			result := algorithm(testCase)
			if !reflect.DeepEqual(result, expected) {
				t.Errorf("Algorithm %T failed for input %v: got %v, expected %v", algorithm, testCase, result, expected)
			}
		}
	}
}
