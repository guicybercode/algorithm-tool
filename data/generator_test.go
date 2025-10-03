package data

import (
	"testing"
)

func TestGenerateArray(t *testing.T) {
	sizes := []int{10, 100, 1000}
	arrayTypes := []ArrayType{Random, Sorted, ReverseSorted}
	
	for _, size := range sizes {
		for _, arrayType := range arrayTypes {
			arr := GenerateArray(size, arrayType)
			
			if len(arr) != size {
				t.Errorf("Expected array length %d, got %d", size, len(arr))
			}
			
			switch arrayType {
			case Sorted:
				if !IsSorted(arr) {
					t.Errorf("Generated array should be sorted for Sorted type")
				}
			case ReverseSorted:
				for i := 1; i < len(arr); i++ {
					if arr[i] > arr[i-1] {
						t.Errorf("Generated array should be reverse sorted for ReverseSorted type")
					}
				}
			}
		}
	}
}

func TestGetArrayTypeName(t *testing.T) {
	tests := []struct {
		arrayType ArrayType
		expected  string
	}{
		{Random, "Random"},
		{Sorted, "Sorted"},
		{ReverseSorted, "Reverse Sorted"},
		{ArrayType(999), "Unknown"},
	}
	
	for _, test := range tests {
		result := GetArrayTypeName(test.arrayType)
		if result != test.expected {
			t.Errorf("GetArrayTypeName(%d) = %s, expected %s", test.arrayType, result, test.expected)
		}
	}
}

func TestGetAllArrayTypes(t *testing.T) {
	arrayTypes := GetAllArrayTypes()
	expected := []ArrayType{Random, Sorted, ReverseSorted}
	
	if len(arrayTypes) != len(expected) {
		t.Errorf("Expected %d array types, got %d", len(expected), len(arrayTypes))
	}
	
	for i, arrayType := range arrayTypes {
		if arrayType != expected[i] {
			t.Errorf("Expected array type %d at index %d, got %d", expected[i], i, arrayType)
		}
	}
}

func TestIsSorted(t *testing.T) {
	tests := []struct {
		arr      []int
		expected bool
	}{
		{[]int{}, true},
		{[]int{1}, true},
		{[]int{1, 2, 3, 4, 5}, true},
		{[]int{1, 1, 1, 1, 1}, true},
		{[]int{5, 4, 3, 2, 1}, false},
		{[]int{1, 3, 2, 4, 5}, false},
		{[]int{1, 2, 3, 2, 4}, false},
	}
	
	for _, test := range tests {
		result := IsSorted(test.arr)
		if result != test.expected {
			t.Errorf("IsSorted(%v) = %t, expected %t", test.arr, result, test.expected)
		}
	}
}

func TestVerifySorting(t *testing.T) {
	tests := []struct {
		original []int
		sorted   []int
		expected bool
	}{
		{[]int{}, []int{}, true},
		{[]int{1}, []int{1}, true},
		{[]int{3, 1, 2}, []int{1, 2, 3}, true},
		{[]int{5, 2, 8, 1, 9}, []int{1, 2, 5, 8, 9}, true},
		{[]int{1, 2, 3}, []int{1, 2, 3}, true},
		{[]int{1, 2, 3}, []int{1, 2, 4}, false},
		{[]int{1, 2, 3}, []int{1, 2}, false},
		{[]int{1, 2}, []int{1, 2, 3}, false},
	}
	
	for _, test := range tests {
		result := VerifySorting(test.original, test.sorted)
		if result != test.expected {
			t.Errorf("VerifySorting(%v, %v) = %t, expected %t", test.original, test.sorted, result, test.expected)
		}
	}
}

func TestGenerateArrayUniqueness(t *testing.T) {
	arr1 := GenerateArray(1000, Random)
	arr2 := GenerateArray(1000, Random)
	
	identical := true
	for i := 0; i < len(arr1) && i < len(arr2); i++ {
		if arr1[i] != arr2[i] {
			identical = false
			break
		}
	}
	
	if identical {
		t.Error("Random arrays should be different between calls")
	}
}

func TestSortedArrayValues(t *testing.T) {
	arr := GenerateArray(100, Sorted)
	
	for i := 0; i < len(arr); i++ {
		if arr[i] != i {
			t.Errorf("Sorted array at index %d should be %d, got %d", i, i, arr[i])
		}
	}
}

func TestReverseSortedArrayValues(t *testing.T) {
	arr := GenerateArray(100, ReverseSorted)
	
	for i := 0; i < len(arr); i++ {
		expected := len(arr) - i - 1
		if arr[i] != expected {
			t.Errorf("Reverse sorted array at index %d should be %d, got %d", i, expected, arr[i])
		}
	}
}
