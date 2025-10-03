package cli

import (
	"algorithm-benchmark/benchmark"
	"algorithm-benchmark/data"
	"algorithm-benchmark/export"
	"flag"
	"fmt"
	"strings"
	"time"
)

type CLI struct {
	benchmarkSuite *benchmark.BenchmarkSuite
}

func NewCLI() *CLI {
	return &CLI{
		benchmarkSuite: benchmark.NewBenchmarkSuite(),
	}
}

func (cli *CLI) Run() {
	cli.RunWithArgs(nil)
}

func (cli *CLI) RunWithArgs(args []string) {
	var (
		algorithm    = flag.String("algorithm", "", "Algorithm to benchmark (linear_search, binary_search, bubble_sort, insertion_sort, merge_sort, quick_sort, heap_sort, native_sort, all)")
		arrayType    = flag.String("array-type", "random", "Array type (random, sorted, reverse)")
		size         = flag.Int("size", 1000, "Array size")
		runs         = flag.Int("runs", 5, "Number of benchmark runs")
		exportCSV    = flag.String("export-csv", "", "Export results to CSV file")
		exportMD     = flag.String("export-md", "", "Export results to Markdown file")
		interactive  = flag.Bool("interactive", false, "Run in interactive mode")
		help         = flag.Bool("help", false, "Show help")
	)
	
	if args != nil {
		flag.CommandLine.Parse(args)
	} else {
		flag.Parse()
	}
	
	if *help {
		cli.showHelp()
		return
	}
	
	if *interactive {
		cli.runInteractive()
		return
	}
	
	if *algorithm == "" {
		fmt.Println("Error: algorithm is required")
		cli.showHelp()
		return
	}
	
	cli.runBenchmark(*algorithm, *arrayType, *size, *runs, *exportCSV, *exportMD)
}

func (cli *CLI) showHelp() {
	fmt.Println("Algorithm Benchmark Tool")
	fmt.Println("Usage: go run main.go [options]")
	fmt.Println("\nOptions:")
	fmt.Println("  -algorithm string")
	fmt.Println("        Algorithm to benchmark (linear_search, binary_search, bubble_sort, insertion_sort, merge_sort, quick_sort, heap_sort, native_sort, all)")
	fmt.Println("  -array-type string")
	fmt.Println("        Array type (random, sorted, reverse) (default \"random\")")
	fmt.Println("  -size int")
	fmt.Println("        Array size (default 1000)")
	fmt.Println("  -runs int")
	fmt.Println("        Number of benchmark runs (default 5)")
	fmt.Println("  -export-csv string")
	fmt.Println("        Export results to CSV file")
	fmt.Println("  -export-md string")
	fmt.Println("        Export results to Markdown file")
	fmt.Println("  -interactive")
	fmt.Println("        Run in interactive mode")
	fmt.Println("  -help")
	fmt.Println("        Show this help message")
	fmt.Println("\nExamples:")
	fmt.Println("  go run main.go -algorithm=quick_sort -size=10000 -runs=10")
	fmt.Println("  go run main.go -algorithm=all -array-type=random -export-csv=results.csv")
	fmt.Println("  go run main.go -interactive")
}

func (cli *CLI) runBenchmark(algorithm, arrayType string, size, runs int, exportCSV, exportMD string) {
	cli.benchmarkSuite.ClearResults()
	
	arrayTypeEnum := cli.parseArrayType(arrayType)
	
	if algorithm == "all" {
		cli.runAllBenchmarks(size, runs)
	} else {
		cli.runSingleBenchmark(algorithm, arrayTypeEnum, size, runs)
	}
	
	cli.displayResults()
	
	if exportCSV != "" {
		if err := export.ExportToCSV(cli.benchmarkSuite.GetResults(), exportCSV); err != nil {
			fmt.Printf("Error exporting to CSV: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", exportCSV)
		}
	}
	
	if exportMD != "" {
		if err := export.ExportToMarkdown(cli.benchmarkSuite.GetResults(), exportMD); err != nil {
			fmt.Printf("Error exporting to Markdown: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", exportMD)
		}
	}
}

func (cli *CLI) runAllBenchmarks(size, runs int) {
	fmt.Println("Running all benchmarks...")
	
	sizes := []int{1000, 10000, 100000, 1000000}
	
	fmt.Println("Running search benchmarks...")
	if err := cli.benchmarkSuite.RunSearchBenchmarks(sizes, runs); err != nil {
		fmt.Printf("Error running search benchmarks: %v\n", err)
		return
	}
	
	fmt.Println("Running sort benchmarks...")
	if err := cli.benchmarkSuite.RunSortBenchmarks(sizes, runs); err != nil {
		fmt.Printf("Error running sort benchmarks: %v\n", err)
		return
	}
}

func (cli *CLI) runSingleBenchmark(algorithm string, arrayType data.ArrayType, size, runs int) {
	config := benchmark.BenchmarkConfig{
		Algorithm: algorithm,
		ArrayType: arrayType,
		Size:      size,
		Runs:      runs,
		Target:    size / 2,
	}
	
	fmt.Printf("Running benchmark for %s with %s array of size %d (%d runs)...\n", 
		algorithm, data.GetArrayTypeName(arrayType), size, runs)
	
	_, err := cli.benchmarkSuite.RunBenchmark(config)
	if err != nil {
		fmt.Printf("Error running benchmark: %v\n", err)
		return
	}
	
	fmt.Printf("Benchmark completed successfully!\n")
}

func (cli *CLI) displayResults() {
	results := cli.benchmarkSuite.GetResults()
	
	if len(results) == 0 {
		fmt.Println("No results to display.")
		return
	}
	
	fmt.Println("\n" + strings.Repeat("=", 80))
	fmt.Println("BENCHMARK RESULTS")
	fmt.Println(strings.Repeat("=", 80))
	
	for _, result := range results {
		fmt.Printf("\nAlgorithm: %s\n", result.Algorithm)
		fmt.Printf("Array Type: %s\n", result.ArrayType)
		fmt.Printf("Size: %d\n", result.Size)
		fmt.Printf("Runs: %d\n", result.Runs)
		fmt.Printf("Mean Duration: %s\n", formatDuration(result.MeanDuration))
		fmt.Printf("Std Deviation: %s\n", formatDuration(result.StdDeviation))
		fmt.Printf("Min Duration: %s\n", formatDuration(result.MinDuration))
		fmt.Printf("Max Duration: %s\n", formatDuration(result.MaxDuration))
		fmt.Printf("Memory Used: %s\n", formatBytes(result.MemoryUsed))
		fmt.Println(strings.Repeat("-", 40))
	}
}

func (cli *CLI) runInteractive() {
	fmt.Println("Algorithm Benchmark Tool - Interactive Mode")
	fmt.Println("==========================================")
	
	for {
		fmt.Println("\nAvailable options:")
		fmt.Println("1. Run single algorithm benchmark")
		fmt.Println("2. Run all search algorithms")
		fmt.Println("3. Run all sort algorithms")
		fmt.Println("4. Run comprehensive benchmark")
		fmt.Println("5. Export results")
		fmt.Println("6. Exit")
		
		var choice int
		fmt.Print("Enter your choice (1-6): ")
		fmt.Scanln(&choice)
		
		switch choice {
		case 1:
			cli.interactiveSingleBenchmark()
		case 2:
			cli.interactiveSearchBenchmarks()
		case 3:
			cli.interactiveSortBenchmarks()
		case 4:
			cli.interactiveComprehensiveBenchmark()
		case 5:
			cli.interactiveExport()
		case 6:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func (cli *CLI) interactiveSingleBenchmark() {
	fmt.Println("\nAvailable algorithms:")
	algorithms := []string{"linear_search", "binary_search", "bubble_sort", "insertion_sort", "merge_sort", "quick_sort", "heap_sort", "native_sort"}
	for i, alg := range algorithms {
		fmt.Printf("%d. %s\n", i+1, alg)
	}
	
	var choice int
	fmt.Print("Select algorithm (1-8): ")
	fmt.Scanln(&choice)
	
	if choice < 1 || choice > 8 {
		fmt.Println("Invalid choice.")
		return
	}
	
	algorithm := algorithms[choice-1]
	
	fmt.Println("\nArray types:")
	fmt.Println("1. Random")
	fmt.Println("2. Sorted")
	fmt.Println("3. Reverse Sorted")
	
	var arrayChoice int
	fmt.Print("Select array type (1-3): ")
	fmt.Scanln(&arrayChoice)
	
	arrayType := data.ArrayType(arrayChoice - 1)
	
	var size int
	fmt.Print("Enter array size: ")
	fmt.Scanln(&size)
	
	var runs int
	fmt.Print("Enter number of runs: ")
	fmt.Scanln(&runs)
	
	cli.benchmarkSuite.ClearResults()
	cli.runSingleBenchmark(algorithm, arrayType, size, runs)
	cli.displayResults()
}

func (cli *CLI) interactiveSearchBenchmarks() {
	var runs int
	fmt.Print("Enter number of runs: ")
	fmt.Scanln(&runs)
	
	cli.benchmarkSuite.ClearResults()
	sizes := []int{1000, 10000, 100000, 1000000}
	
	fmt.Println("Running search benchmarks...")
	if err := cli.benchmarkSuite.RunSearchBenchmarks(sizes, runs); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	cli.displayResults()
}

func (cli *CLI) interactiveSortBenchmarks() {
	var runs int
	fmt.Print("Enter number of runs: ")
	fmt.Scanln(&runs)
	
	cli.benchmarkSuite.ClearResults()
	sizes := []int{1000, 10000, 100000, 1000000}
	
	fmt.Println("Running sort benchmarks...")
	if err := cli.benchmarkSuite.RunSortBenchmarks(sizes, runs); err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	
	cli.displayResults()
}

func (cli *CLI) interactiveComprehensiveBenchmark() {
	var runs int
	fmt.Print("Enter number of runs: ")
	fmt.Scanln(&runs)
	
	cli.runAllBenchmarks(0, runs)
	cli.displayResults()
}

func (cli *CLI) interactiveExport() {
	results := cli.benchmarkSuite.GetResults()
	if len(results) == 0 {
		fmt.Println("No results to export.")
		return
	}
	
	fmt.Println("\nExport options:")
	fmt.Println("1. Export to CSV")
	fmt.Println("2. Export to Markdown")
	fmt.Println("3. Export to both")
	
	var choice int
	fmt.Print("Select export format (1-3): ")
	fmt.Scanln(&choice)
	
	timestamp := time.Now().Format("20060102_150405")
	
	switch choice {
	case 1:
		filename := fmt.Sprintf("results_%s.csv", timestamp)
		if err := export.ExportToCSV(results, filename); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", filename)
		}
	case 2:
		filename := fmt.Sprintf("results_%s.md", timestamp)
		if err := export.ExportToMarkdown(results, filename); err != nil {
			fmt.Printf("Error: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", filename)
		}
	case 3:
		csvFile := fmt.Sprintf("results_%s.csv", timestamp)
		mdFile := fmt.Sprintf("results_%s.md", timestamp)
		
		if err := export.ExportToCSV(results, csvFile); err != nil {
			fmt.Printf("Error exporting CSV: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", csvFile)
		}
		
		if err := export.ExportToMarkdown(results, mdFile); err != nil {
			fmt.Printf("Error exporting Markdown: %v\n", err)
		} else {
			fmt.Printf("Results exported to %s\n", mdFile)
		}
	default:
		fmt.Println("Invalid choice.")
	}
}

func (cli *CLI) parseArrayType(arrayType string) data.ArrayType {
	switch strings.ToLower(arrayType) {
	case "sorted":
		return data.Sorted
	case "reverse":
		return data.ReverseSorted
	default:
		return data.Random
	}
}

func formatDuration(d time.Duration) string {
	if d < time.Microsecond {
		return fmt.Sprintf("%.2f ns", float64(d.Nanoseconds()))
	} else if d < time.Millisecond {
		return fmt.Sprintf("%.2f μs", float64(d.Nanoseconds())/1000)
	} else if d < time.Second {
		return fmt.Sprintf("%.2f ms", float64(d.Nanoseconds())/1000000)
	} else {
		return fmt.Sprintf("%.2f s", d.Seconds())
	}
}

func formatBytes(bytes uint64) string {
	const unit = 1024
	if bytes < unit {
		return fmt.Sprintf("%d B", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(bytes)/float64(div), "KMGTPE"[exp])
}
