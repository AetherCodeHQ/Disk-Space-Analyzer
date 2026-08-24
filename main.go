
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	dir := "."
	if len(os.Args) > 1 {
		dir = os.Args[1]
	}
	files := 0
	lines := 0
	var sizes []int64
	filepath.Walk(dir, func(p string, info os.FileInfo, err error) error {
		if err != nil || info.IsDir() {
			return nil
		}
		if strings.HasPrefix(info.Name(), ".") {
			return nil
		}
		files++
		lines += countLines(p)
		sizes = append(sizes, info.Size())
		return nil
	})
	fmt.Printf("files=%d lines=%d size=%d\n", files, lines, sum(sizes))
}

func countLines(p string) int {
	b, err := ioutil.ReadFile(p)
	if err != nil {
		return 0
	}
	return strings.Count(string(b), "\n") + 1
}

func sum(xs []int64) int64 {
	var s int64
	for _, x := range xs {
		s += x
	}
	return s
}
