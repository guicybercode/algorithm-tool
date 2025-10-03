# Algorithm Benchmark Tool

A comprehensive Go application for benchmarking search and sorting algorithms with both command-line and web interfaces. This tool measures execution time and memory usage of various algorithms across different input sizes and data patterns.

## Features

### Algorithms Implemented

**Search Algorithms:**
- Linear Search
- Binary Search (for both sorted and unsorted arrays)

**Sorting Algorithms:**
- Bubble Sort
- Insertion Sort
- Merge Sort
- Quick Sort
- Heap Sort
- Native Sort (Go's built-in sort.Ints)

### Benchmark Capabilities

- **Multiple Input Sizes:** 1,000, 10,000, 100,000, and 1,000,000 elements
- **Data Patterns:** Random, sorted, and reverse-sorted arrays
- **Performance Metrics:** Execution time, memory usage, statistical analysis
- **Multiple Runs:** Configurable number of benchmark runs for statistical accuracy
- **Export Options:** CSV and Markdown report generation

### Interfaces

1. **Command Line Interface (CLI):** Full-featured terminal interface with interactive mode
2. **Web Interface:** Modern web UI with real-time charts and interactive controls

## Installation

### Prerequisites

- Go 1.21 or later
- Git (for cloning the repository)

### Setup

1. Clone the repository:
```bash
git clone <repository-url>
cd algorithm-benchmark
```

2. Install dependencies:
```bash
go mod tidy
```

3. Build the application:
```bash
go build -o benchmark main.go
```

## Usage

### Command Line Interface

#### Basic Usage

```bash
# Run a single algorithm benchmark
go run main.go -algorithm=quick_sort -size=10000 -runs=5

# Run all algorithms with comprehensive testing
go run main.go -algorithm=all -array-type=random -runs=3

# Interactive mode
go run main.go -interactive

# Export results
go run main.go -algorithm=all -export-csv=results.csv -export-md=results.md
```

#### CLI Options

- `-algorithm`: Algorithm to benchmark (linear_search, binary_search, bubble_sort, insertion_sort, merge_sort, quick_sort, heap_sort, native_sort, all)
- `-array-type`: Array type (random, sorted, reverse)
- `-size`: Array size (default: 1000)
- `-runs`: Number of benchmark runs (default: 5)
- `-export-csv`: Export results to CSV file
- `-export-md`: Export results to Markdown file
- `-interactive`: Run in interactive mode
- `-help`: Show help message

#### Examples

```bash
# Benchmark Quick Sort with random array of 50,000 elements
go run main.go -algorithm=quick_sort -array-type=random -size=50000 -runs=10

# Compare all sorting algorithms with sorted input
go run main.go -algorithm=all -array-type=sorted -export-csv=sort_results.csv

# Interactive mode for guided benchmarking
go run main.go -interactive
```

### Web Interface

#### Starting the Web Server

```bash
# Start web server on default port 8080
go run main.go -mode=web

# Start web server on custom port
go run main.go -mode=web -port=9090
```

Then open your browser and navigate to `http://localhost:8080` (or your custom port).

#### Web Interface Features

- **Single Algorithm Benchmarking:** Select algorithm, array type, size, and number of runs
- **Comprehensive Benchmarking:** Run all algorithms with multiple sizes and array types
- **Real-time Results:** View results in tabular format with performance metrics
- **Interactive Charts:** Visualize performance comparisons with Chart.js
- **Export Functionality:** Download results as CSV or Markdown files
- **Result Management:** Clear results and manage multiple benchmark sessions

## Project Structure

```
algorithm-benchmark/
├── algorithms/          # Algorithm implementations
│   ├── search.go       # Search algorithms
│   ├── sort.go         # Sorting algorithms
│   └── *_test.go       # Algorithm tests
├── benchmark/          # Benchmarking framework
│   ├── benchmark.go    # Core benchmarking logic
│   └── benchmark_test.go
├── cli/                # Command-line interface
│   └── cli.go
├── data/               # Data generation utilities
│   ├── generator.go    # Array generation functions
│   └── generator_test.go
├── export/             # Export functionality
│   └── export.go       # CSV and Markdown export
├── web/                # Web interface
│   ├── web.go          # Web server implementation
│   └── templates/      # HTML templates
│       └── index.html  # Main web interface
├── main.go             # Application entry point
├── go.mod              # Go module definition
└── README.md           # This file
```

## Testing

Run the test suite:

```bash
# Run all tests
go test ./...

# Run tests with verbose output
go test -v ./...

# Run benchmarks
go test -bench=. ./...

# Run tests with coverage
go test -cover ./...
```

## Performance Analysis

### Benchmark Results Interpretation

The tool provides comprehensive performance metrics:

- **Mean Duration:** Average execution time across multiple runs
- **Standard Deviation:** Statistical measure of performance consistency
- **Min/Max Duration:** Best and worst case execution times
- **Memory Usage:** Memory consumption during algorithm execution
- **Statistical Analysis:** Multiple runs provide reliable performance data

### Expected Performance Characteristics

**Search Algorithms:**
- Linear Search: O(n) time complexity
- Binary Search: O(log n) time complexity (requires sorted input)

**Sorting Algorithms:**
- Bubble Sort: O(n²) time complexity
- Insertion Sort: O(n²) time complexity, good for small datasets
- Merge Sort: O(n log n) time complexity, stable sorting
- Quick Sort: O(n log n) average case, O(n²) worst case
- Heap Sort: O(n log n) time complexity
- Native Sort: Go's optimized implementation

## Export Formats

### CSV Export
Results are exported in CSV format with the following columns:
- Algorithm
- Array Type
- Size
- Mean Duration (nanoseconds)
- Standard Deviation (nanoseconds)
- Min/Max Duration (nanoseconds)
- Memory Used (bytes)
- Number of Runs

### Markdown Export
Comprehensive reports in Markdown format including:
- Summary tables
- Detailed results by algorithm
- Performance comparisons
- Statistical analysis

## Advanced Usage

### Custom Benchmarking

For programmatic usage, you can use the benchmark package directly:

```go
package main

import (
    "algorithm-benchmark/benchmark"
    "algorithm-benchmark/data"
    "fmt"
)

func main() {
    suite := benchmark.NewBenchmarkSuite()
    
    config := benchmark.BenchmarkConfig{
        Algorithm: "quick_sort",
        ArrayType: data.Random,
        Size:      10000,
        Runs:      5,
        Target:    0,
    }
    
    result, err := suite.RunBenchmark(config)
    if err != nil {
        fmt.Printf("Error: %v\n", err)
        return
    }
    
    fmt.Printf("Mean duration: %v\n", result.MeanDuration)
}
```

### Parallel Benchmarking

The application supports running multiple benchmarks in parallel using Go's goroutines. This is particularly useful for comprehensive testing across multiple algorithms and input sizes.

## Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Ensure all tests pass
6. Submit a pull request

## License

This project is open source and available under the MIT License.

## Performance Notes

- For large datasets (>100,000 elements), some algorithms may take significant time
- Memory usage scales with input size
- Consider system resources when running comprehensive benchmarks
- Web interface may become unresponsive during large benchmark runs

## Troubleshooting

### Common Issues

1. **Out of Memory:** Reduce array size or number of runs
2. **Slow Performance:** Use smaller input sizes for initial testing
3. **Web Interface Not Loading:** Check if port is available and firewall settings
4. **Export Failures:** Ensure write permissions in the target directory

### Performance Optimization

- Use `-runs=1` for quick testing
- Start with smaller array sizes (1000-10000)
- Use specific algorithms instead of "all" for faster execution
- Monitor system resources during comprehensive benchmarks
