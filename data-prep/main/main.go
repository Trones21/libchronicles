package main

import (
	"fmt"
	combinejson "libchronicles/combine-json"
	getdata "libchronicles/get-data"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run main.go <command>")
		fmt.Println("Available commands:")
		fmt.Println("  get-data      - Extract metadata from the Go repository")
		fmt.Println("  combine-json  - Combine JSON files into a single output")
		return
	}

	command := os.Args[1]

	switch command {
	case "get-data":
		fmt.Println("Running get-data...")
		getdata.Run()
	case "combine-json":
		fmt.Println("Running combine-json...")
		combinejson.Run()
	default:
		fmt.Printf("Unknown command: %s\n", command)
		fmt.Println("Available commands:")
		fmt.Println("  get-data      - Extract metadata from the Go repository")
		fmt.Println("  combine-json  - Combine JSON files into a single output")
	}
}
