package data

import (
	"math/rand"
	"sort"
	"time"
)

type ArrayType int

const (
	Random ArrayType = iota
	Sorted
	ReverseSorted
)

func GenerateArray(size int, arrayType ArrayType) []int {
	rand.Seed(time.Now().UnixNano())
	
	arr := make([]int, size)
	
	switch arrayType {
	case Random:
		for i := 0; i < size; i++ {
			arr[i] = rand.Intn(size * 2)
		}
	case Sorted:
		for i := 0; i < size; i++ {
			arr[i] = i
		}
	case ReverseSorted:
		for i := 0; i < size; i++ {
			arr[i] = size - i - 1
		}
	}
	
	return arr
}

func GetArrayTypeName(arrayType ArrayType) string {
	switch arrayType {
	case Random:
		return "Random"
	case Sorted:
		return "Sorted"
	case ReverseSorted:
		return "Reverse Sorted"
	default:
		return "Unknown"
	}
}

func GetAllArrayTypes() []ArrayType {
	return []ArrayType{Random, Sorted, ReverseSorted}
}

func IsSorted(arr []int) bool {
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			return false
		}
	}
	return true
}

func VerifySorting(original, sorted []int) bool {
	if len(original) != len(sorted) {
		return false
	}
	
	originalCopy := make([]int, len(original))
	copy(originalCopy, original)
	sort.Ints(originalCopy)
	
	for i := 0; i < len(originalCopy); i++ {
		if originalCopy[i] != sorted[i] {
			return false
		}
	}
	return true
}
