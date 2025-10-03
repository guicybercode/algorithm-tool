package export

import (
	"algorithm-benchmark/benchmark"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ExportToCSV(results []benchmark.BenchmarkResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	writer := csv.NewWriter(file)
	defer writer.Flush()
	
	header := []string{
		"Algorithm",
		"Array Type",
		"Size",
		"Mean Duration (ns)",
		"Std Deviation (ns)",
		"Min Duration (ns)",
		"Max Duration (ns)",
		"Memory Used (bytes)",
		"Runs",
	}
	
	if err := writer.Write(header); err != nil {
		return err
	}
	
	for _, result := range results {
		record := []string{
			result.Algorithm,
			result.ArrayType,
			strconv.Itoa(result.Size),
			strconv.FormatInt(result.MeanDuration.Nanoseconds(), 10),
			strconv.FormatInt(result.StdDeviation.Nanoseconds(), 10),
			strconv.FormatInt(result.MinDuration.Nanoseconds(), 10),
			strconv.FormatInt(result.MaxDuration.Nanoseconds(), 10),
			strconv.FormatUint(result.MemoryUsed, 10),
			strconv.Itoa(result.Runs),
		}
		
		if err := writer.Write(record); err != nil {
			return err
		}
	}
	
	return nil
}

func ExportToMarkdown(results []benchmark.BenchmarkResult, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()
	
	content := generateMarkdownContent(results)
	_, err = file.WriteString(content)
	return err
}

func generateMarkdownContent(results []benchmark.BenchmarkResult) string {
	var sb strings.Builder
	
	sb.WriteString("# Algorithm Benchmark Results\n\n")
	sb.WriteString(fmt.Sprintf("Generated on: %s\n\n", time.Now().Format("2006-01-02 15:04:05")))
	
	sb.WriteString("## Summary\n\n")
	sb.WriteString("| Algorithm | Array Type | Size | Mean Duration | Std Deviation | Memory Used | Runs |\n")
	sb.WriteString("|-----------|------------|------|---------------|---------------|-------------|------|\n")
	
	for _, result := range results {
		sb.WriteString(fmt.Sprintf("| %s | %s | %d | %s | %s | %s | %d |\n",
			result.Algorithm,
			result.ArrayType,
			result.Size,
			formatDuration(result.MeanDuration),
			formatDuration(result.StdDeviation),
			formatBytes(result.MemoryUsed),
			result.Runs,
		))
	}
	
	sb.WriteString("\n## Detailed Results\n\n")
	
	algorithms := getUniqueAlgorithms(results)
	for _, algorithm := range algorithms {
		sb.WriteString(fmt.Sprintf("### %s\n\n", algorithm))
		sb.WriteString("| Array Type | Size | Mean | Std Dev | Min | Max | Memory |\n")
		sb.WriteString("|------------|------|------|---------|-----|-----|--------|\n")
		
		algorithmResults := filterByAlgorithm(results, algorithm)
		for _, result := range algorithmResults {
			sb.WriteString(fmt.Sprintf("| %s | %d | %s | %s | %s | %s | %s |\n",
				result.ArrayType,
				result.Size,
				formatDuration(result.MeanDuration),
				formatDuration(result.StdDeviation),
				formatDuration(result.MinDuration),
				formatDuration(result.MaxDuration),
				formatBytes(result.MemoryUsed),
			))
		}
		sb.WriteString("\n")
	}
	
	return sb.String()
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

func getUniqueAlgorithms(results []benchmark.BenchmarkResult) []string {
	algorithms := make(map[string]bool)
	for _, result := range results {
		algorithms[result.Algorithm] = true
	}
	
	var unique []string
	for algorithm := range algorithms {
		unique = append(unique, algorithm)
	}
	return unique
}

func filterByAlgorithm(results []benchmark.BenchmarkResult, algorithm string) []benchmark.BenchmarkResult {
	var filtered []benchmark.BenchmarkResult
	for _, result := range results {
		if result.Algorithm == algorithm {
			filtered = append(filtered, result)
		}
	}
	return filtered
}
