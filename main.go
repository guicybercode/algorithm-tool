package main

import (
	"algorithm-benchmark/cli"
	"algorithm-benchmark/web"
	"fmt"
	"os"
	"strings"
)

func main() {
	args := os.Args[1:]
	
	// Check for mode flag first
	mode := "cli"
	port := "8080"
	
	for i, arg := range args {
		if arg == "-mode" && i+1 < len(args) {
			mode = args[i+1]
		} else if arg == "-port" && i+1 < len(args) {
			port = args[i+1]
		} else if strings.HasPrefix(arg, "-mode=") {
			mode = strings.TrimPrefix(arg, "-mode=")
		} else if strings.HasPrefix(arg, "-port=") {
			port = strings.TrimPrefix(arg, "-port=")
		}
	}
	
	// Remove mode and port flags from args for CLI
	if mode == "cli" {
		filteredArgs := make([]string, 0)
		for i, arg := range args {
			if arg == "-mode" || arg == "-port" {
				// Skip the flag and its value
				continue
			} else if strings.HasPrefix(arg, "-mode=") || strings.HasPrefix(arg, "-port=") {
				// Skip the flag with value
				continue
			} else if i > 0 && (args[i-1] == "-mode" || args[i-1] == "-port") {
				// Skip the value of the previous flag
				continue
			} else {
				filteredArgs = append(filteredArgs, arg)
			}
		}
		args = filteredArgs
	}

	switch mode {
	case "cli":
		cliApp := cli.NewCLI()
		cliApp.RunWithArgs(args)
	case "web":
		webServer := web.NewWebServer()
		webServer.Start(port)
	default:
		fmt.Printf("Invalid mode: %s. Use 'cli' or 'web'\n", mode)
		fmt.Println("Usage:")
		fmt.Println("  go run main.go -mode=cli    # Command line interface")
		fmt.Println("  go run main.go -mode=web    # Web interface")
		fmt.Println("  go run main.go -mode=web -port=9090  # Web interface on port 9090")
		os.Exit(1)
	}
}
