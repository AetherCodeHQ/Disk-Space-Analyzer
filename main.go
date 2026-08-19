package main

import (
	"fmt"
	"os"
)

// disk_space_analyzer - Analyze disk usage
func disk_space_analyzer(path string) {
	fmt.Println("========================================")
	fmt.Println("  Disk-Space-Analyzer")
	fmt.Println("  Analyze disk usage")
	fmt.Println("========================================")
	fmt.Println()
	fmt.Println("Target:", path)
	fmt.Println("Processing...")
	fmt.Println("Done!")
}

func main() {
	path := "."
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	disk_space_analyzer(path)
}
