package web

import (
	"algorithm-benchmark/benchmark"
	"algorithm-benchmark/data"
	"algorithm-benchmark/export"
	"encoding/json"
	"fmt"
	"html/template"
	"net/http"
	"time"
)

type WebServer struct {
	benchmarkSuite *benchmark.BenchmarkSuite
	templates      *template.Template
}

type BenchmarkRequest struct {
	Algorithm string `json:"algorithm"`
	ArrayType string `json:"arrayType"`
	Size      int    `json:"size"`
	Runs      int    `json:"runs"`
}

type BenchmarkResponse struct {
	Success bool                     `json:"success"`
	Message string                   `json:"message,omitempty"`
	Results []benchmark.BenchmarkResult `json:"results,omitempty"`
}

func NewWebServer() *WebServer {
	templates := template.Must(template.ParseGlob("web/templates/*.html"))
	
	return &WebServer{
		benchmarkSuite: benchmark.NewBenchmarkSuite(),
		templates:      templates,
	}
}

func (ws *WebServer) Start(port string) {
	http.HandleFunc("/", ws.handleIndex)
	http.HandleFunc("/api/benchmark", ws.handleBenchmark)
	http.HandleFunc("/api/benchmark/all", ws.handleBenchmarkAll)
	http.HandleFunc("/api/export/csv", ws.handleExportCSV)
	http.HandleFunc("/api/export/md", ws.handleExportMarkdown)
	http.HandleFunc("/api/results", ws.handleGetResults)
	http.HandleFunc("/api/clear", ws.handleClearResults)
	
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static/"))))
	
	fmt.Printf("Web server starting on port %s\n", port)
	fmt.Printf("Open your browser and go to: http://localhost:%s\n", port)
	
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		fmt.Printf("Error starting server: %v\n", err)
	}
}

func (ws *WebServer) handleIndex(w http.ResponseWriter, r *http.Request) {
	if err := ws.templates.ExecuteTemplate(w, "index.html", nil); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (ws *WebServer) handleBenchmark(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req BenchmarkRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ws.sendJSONResponse(w, BenchmarkResponse{
			Success: false,
			Message: "Invalid JSON request",
		}, http.StatusBadRequest)
		return
	}
	
	arrayType := ws.parseArrayType(req.ArrayType)
	
	config := benchmark.BenchmarkConfig{
		Algorithm: req.Algorithm,
		ArrayType: arrayType,
		Size:      req.Size,
		Runs:      req.Runs,
		Target:    req.Size / 2,
	}
	
	result, err := ws.benchmarkSuite.RunBenchmark(config)
	if err != nil {
		ws.sendJSONResponse(w, BenchmarkResponse{
			Success: false,
			Message: fmt.Sprintf("Benchmark failed: %v", err),
		}, http.StatusInternalServerError)
		return
	}
	
	ws.sendJSONResponse(w, BenchmarkResponse{
		Success: true,
		Results: []benchmark.BenchmarkResult{result},
	}, http.StatusOK)
}

func (ws *WebServer) handleBenchmarkAll(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	var req struct {
		Runs int `json:"runs"`
	}
	
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		ws.sendJSONResponse(w, BenchmarkResponse{
			Success: false,
			Message: "Invalid JSON request",
		}, http.StatusBadRequest)
		return
	}
	
	ws.benchmarkSuite.ClearResults()
	
	sizes := []int{1000, 10000, 100000, 1000000}
	
	if err := ws.benchmarkSuite.RunSearchBenchmarks(sizes, req.Runs); err != nil {
		ws.sendJSONResponse(w, BenchmarkResponse{
			Success: false,
			Message: fmt.Sprintf("Search benchmarks failed: %v", err),
		}, http.StatusInternalServerError)
		return
	}
	
	if err := ws.benchmarkSuite.RunSortBenchmarks(sizes, req.Runs); err != nil {
		ws.sendJSONResponse(w, BenchmarkResponse{
			Success: false,
			Message: fmt.Sprintf("Sort benchmarks failed: %v", err),
		}, http.StatusInternalServerError)
		return
	}
	
	ws.sendJSONResponse(w, BenchmarkResponse{
		Success: true,
		Results: ws.benchmarkSuite.GetResults(),
	}, http.StatusOK)
}

func (ws *WebServer) handleGetResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	ws.sendJSONResponse(w, BenchmarkResponse{
		Success: true,
		Results: ws.benchmarkSuite.GetResults(),
	}, http.StatusOK)
}

func (ws *WebServer) handleClearResults(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	ws.benchmarkSuite.ClearResults()
	
	ws.sendJSONResponse(w, BenchmarkResponse{
		Success: true,
		Message: "Results cleared",
	}, http.StatusOK)
}

func (ws *WebServer) handleExportCSV(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	results := ws.benchmarkSuite.GetResults()
	if len(results) == 0 {
		http.Error(w, "No results to export", http.StatusBadRequest)
		return
	}
	
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("benchmark_results_%s.csv", timestamp)
	
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "text/csv")
	
	if err := export.ExportToCSV(results, ""); err != nil {
		http.Error(w, fmt.Sprintf("Export failed: %v", err), http.StatusInternalServerError)
		return
	}
}

func (ws *WebServer) handleExportMarkdown(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	
	results := ws.benchmarkSuite.GetResults()
	if len(results) == 0 {
		http.Error(w, "No results to export", http.StatusBadRequest)
		return
	}
	
	timestamp := time.Now().Format("20060102_150405")
	filename := fmt.Sprintf("benchmark_results_%s.md", timestamp)
	
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	w.Header().Set("Content-Type", "text/markdown")
	
	if err := export.ExportToMarkdown(results, ""); err != nil {
		http.Error(w, fmt.Sprintf("Export failed: %v", err), http.StatusInternalServerError)
		return
	}
}

func (ws *WebServer) parseArrayType(arrayType string) data.ArrayType {
	switch arrayType {
	case "sorted":
		return data.Sorted
	case "reverse":
		return data.ReverseSorted
	default:
		return data.Random
	}
}

func (ws *WebServer) sendJSONResponse(w http.ResponseWriter, response BenchmarkResponse, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}
