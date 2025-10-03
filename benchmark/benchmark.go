package benchmark

import (
	"algorithm-benchmark/algorithms"
	"algorithm-benchmark/data"
	"fmt"
	"runtime"
	"time"
)

type BenchmarkResult struct {
	Algorithm     string
	ArrayType     string
	Size          int
	Duration      time.Duration
	MemoryBefore  uint64
	MemoryAfter   uint64
	MemoryUsed    uint64
	Runs          int
	MeanDuration  time.Duration
	StdDeviation  time.Duration
	MinDuration   time.Duration
	MaxDuration   time.Duration
}

type BenchmarkConfig struct {
	Algorithm string
	ArrayType data.ArrayType
	Size      int
	Runs      int
	Target    int
}

type BenchmarkSuite struct {
	results []BenchmarkResult
}

func NewBenchmarkSuite() *BenchmarkSuite {
	return &BenchmarkSuite{
		results: make([]BenchmarkResult, 0),
	}
}

func (bs *BenchmarkSuite) RunBenchmark(config BenchmarkConfig) (BenchmarkResult, error) {
	var durations []time.Duration
	var memoryUsages []uint64
	
	for run := 0; run < config.Runs; run++ {
		arr := data.GenerateArray(config.Size, config.ArrayType)
		
		var memStatsBefore, memStatsAfter runtime.MemStats
		runtime.GC()
		runtime.ReadMemStats(&memStatsBefore)
		
		start := time.Now()
		
		var result interface{}
		var err error
		
		switch config.Algorithm {
		case "linear_search":
			result = algorithms.LinearSearch(arr, config.Target)
		case "binary_search":
			if config.ArrayType == data.Sorted {
				result = algorithms.BinarySearchSorted(arr, config.Target)
			} else {
				result = algorithms.BinarySearchUnsorted(arr, config.Target)
			}
		case "bubble_sort":
			result = algorithms.BubbleSort(arr)
		case "insertion_sort":
			result = algorithms.InsertionSort(arr)
		case "merge_sort":
			result = algorithms.MergeSort(arr)
		case "quick_sort":
			result = algorithms.QuickSort(arr)
		case "heap_sort":
			result = algorithms.HeapSort(arr)
		case "native_sort":
			result = algorithms.NativeSort(arr)
		default:
			return BenchmarkResult{}, fmt.Errorf("unknown algorithm: %s", config.Algorithm)
		}
		
		duration := time.Since(start)
		
		runtime.ReadMemStats(&memStatsAfter)
		memoryUsed := memStatsAfter.Alloc - memStatsBefore.Alloc
		
		if err != nil {
			return BenchmarkResult{}, err
		}
		
		durations = append(durations, duration)
		memoryUsages = append(memoryUsages, memoryUsed)
		
		if config.Algorithm != "linear_search" && config.Algorithm != "binary_search" {
			if sortedResult, ok := result.([]int); ok {
				if !data.VerifySorting(arr, sortedResult) {
					return BenchmarkResult{}, fmt.Errorf("sorting verification failed for %s", config.Algorithm)
				}
			}
		}
	}
	
	meanDuration := calculateMean(durations)
	stdDeviation := calculateStdDeviation(durations, meanDuration)
	minDuration := calculateMin(durations)
	maxDuration := calculateMax(durations)
	
	meanMemory := calculateMeanUint64(memoryUsages)
	
	benchmarkResult := BenchmarkResult{
		Algorithm:     config.Algorithm,
		ArrayType:     data.GetArrayTypeName(config.ArrayType),
		Size:          config.Size,
		Duration:      meanDuration,
		MemoryUsed:    meanMemory,
		Runs:          config.Runs,
		MeanDuration:  meanDuration,
		StdDeviation:  stdDeviation,
		MinDuration:   minDuration,
		MaxDuration:   maxDuration,
	}
	
	bs.results = append(bs.results, benchmarkResult)
	return benchmarkResult, nil
}

func (bs *BenchmarkSuite) RunSearchBenchmarks(sizes []int, runs int) error {
	searchAlgorithms := []string{"linear_search", "binary_search"}
	arrayTypes := data.GetAllArrayTypes()
	
	for _, algorithm := range searchAlgorithms {
		for _, arrayType := range arrayTypes {
			for _, size := range sizes {
				config := BenchmarkConfig{
					Algorithm: algorithm,
					ArrayType: arrayType,
					Size:      size,
					Runs:      runs,
					Target:    size / 2,
				}
				
				_, err := bs.RunBenchmark(config)
				if err != nil {
					return fmt.Errorf("benchmark failed for %s: %v", algorithm, err)
				}
			}
		}
	}
	
	return nil
}

func (bs *BenchmarkSuite) RunSortBenchmarks(sizes []int, runs int) error {
	sortAlgorithms := []string{"bubble_sort", "insertion_sort", "merge_sort", "quick_sort", "heap_sort", "native_sort"}
	arrayTypes := data.GetAllArrayTypes()
	
	for _, algorithm := range sortAlgorithms {
		for _, arrayType := range arrayTypes {
			for _, size := range sizes {
				config := BenchmarkConfig{
					Algorithm: algorithm,
					ArrayType: arrayType,
					Size:      size,
					Runs:      runs,
					Target:    0,
				}
				
				_, err := bs.RunBenchmark(config)
				if err != nil {
					return fmt.Errorf("benchmark failed for %s: %v", algorithm, err)
				}
			}
		}
	}
	
	return nil
}

func (bs *BenchmarkSuite) GetResults() []BenchmarkResult {
	return bs.results
}

func (bs *BenchmarkSuite) ClearResults() {
	bs.results = make([]BenchmarkResult, 0)
}

func calculateMean(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	
	var total time.Duration
	for _, d := range durations {
		total += d
	}
	return total / time.Duration(len(durations))
}

func calculateStdDeviation(durations []time.Duration, mean time.Duration) time.Duration {
	if len(durations) <= 1 {
		return 0
	}
	
	var sumSquares time.Duration
	for _, d := range durations {
		diff := d - mean
		sumSquares += diff * diff
	}
	
	variance := sumSquares / time.Duration(len(durations)-1)
	return time.Duration(int64(variance) / int64(time.Microsecond)) * time.Microsecond
}

func calculateMin(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	
	min := durations[0]
	for _, d := range durations {
		if d < min {
			min = d
		}
	}
	return min
}

func calculateMax(durations []time.Duration) time.Duration {
	if len(durations) == 0 {
		return 0
	}
	
	max := durations[0]
	for _, d := range durations {
		if d > max {
			max = d
		}
	}
	return max
}

func calculateMeanUint64(values []uint64) uint64 {
	if len(values) == 0 {
		return 0
	}
	
	var total uint64
	for _, v := range values {
		total += v
	}
	return total / uint64(len(values))
}
