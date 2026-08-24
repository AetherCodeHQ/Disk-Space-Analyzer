package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sort"
)

type Entry struct {
	Path string
	Size int64
}

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	var entries []Entry
	filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err == nil && !info.IsDir() {
			entries = append(entries, Entry{path, info.Size()})
		}
		return nil
	})
	sort.Slice(entries, func(i, j int) bool {
		return entries[i].Size > entries[j].Size
	})
	fmt.Println("Disk Space Analyzer")
	fmt.Println("===================")
	var total int64
	fmt.Println("Top 10 largest files:")
	for i, e := range entries {
		if i >= 10 {
			break
		}
		total += e.Size
		fmt.Printf("  %-50s %s\n", e.Path, humanSize(e.Size))
	}
	fmt.Printf("\nTotal scanned: %s\n", humanSize(total))
	fmt.Printf("Files: %d\n", len(entries))
}

func humanSize(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %cB", float64(b)/float64(div), "KMGTPE"[exp])
}