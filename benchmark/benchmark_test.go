package benchmark

import (
	"algorithm-benchmark/data"
	"testing"
)

func TestBenchmarkSuite(t *testing.T) {
	suite := NewBenchmarkSuite()
	
	config := BenchmarkConfig{
		Algorithm: "bubble_sort",
		ArrayType: data.Random,
		Size:      100,
		Runs:      3,
		Target:    0,
	}
	
	result, err := suite.RunBenchmark(config)
	if err != nil {
		t.Fatalf("Benchmark failed: %v", err)
	}
	
	if result.Algorithm != "bubble_sort" {
		t.Errorf("Expected algorithm 'bubble_sort', got '%s'", result.Algorithm)
	}
	
	if result.Size != 100 {
		t.Errorf("Expected size 100, got %d", result.Size)
	}
	
	if result.Runs != 3 {
		t.Errorf("Expected runs 3, got %d", result.Runs)
	}
	
	if result.MeanDuration <= 0 {
		t.Error("Expected positive mean duration")
	}
}

func TestSearchBenchmarks(t *testing.T) {
	suite := NewBenchmarkSuite()
	
	sizes := []int{100, 1000}
	runs := 2
	
	err := suite.RunSearchBenchmarks(sizes, runs)
	if err != nil {
		t.Fatalf("Search benchmarks failed: %v", err)
	}
	
	results := suite.GetResults()
	if len(results) == 0 {
		t.Error("Expected results from search benchmarks")
	}
	
	searchAlgorithms := []string{"linear_search", "binary_search"}
	algorithmCount := make(map[string]int)
	
	for _, result := range results {
		algorithmCount[result.Algorithm]++
	}
	
	for _, algorithm := range searchAlgorithms {
		if algorithmCount[algorithm] == 0 {
			t.Errorf("Expected results for algorithm %s", algorithm)
		}
	}
}

func TestSortBenchmarks(t *testing.T) {
	suite := NewBenchmarkSuite()
	
	sizes := []int{100, 1000}
	runs := 2
	
	err := suite.RunSortBenchmarks(sizes, runs)
	if err != nil {
		t.Fatalf("Sort benchmarks failed: %v", err)
	}
	
	results := suite.GetResults()
	if len(results) == 0 {
		t.Error("Expected results from sort benchmarks")
	}
	
	sortAlgorithms := []string{"bubble_sort", "insertion_sort", "merge_sort", "quick_sort", "heap_sort", "native_sort"}
	algorithmCount := make(map[string]int)
	
	for _, result := range results {
		algorithmCount[result.Algorithm]++
	}
	
	for _, algorithm := range sortAlgorithms {
		if algorithmCount[algorithm] == 0 {
			t.Errorf("Expected results for algorithm %s", algorithm)
		}
	}
}

func TestClearResults(t *testing.T) {
	suite := NewBenchmarkSuite()
	
	config := BenchmarkConfig{
		Algorithm: "bubble_sort",
		ArrayType: data.Random,
		Size:      100,
		Runs:      1,
		Target:    0,
	}
	
	_, err := suite.RunBenchmark(config)
	if err != nil {
		t.Fatalf("Benchmark failed: %v", err)
	}
	
	if len(suite.GetResults()) == 0 {
		t.Error("Expected results after benchmark")
	}
	
	suite.ClearResults()
	
	if len(suite.GetResults()) != 0 {
		t.Error("Expected empty results after clear")
	}
}
