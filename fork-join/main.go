package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CodeDepth struct {
	file  string
	level int
}

func main() {
	dir := os.Args[1]
	partialResults := make(chan CodeDepth)
	wg := sync.WaitGroup{}
	filepath.Walk(dir,
		func(path string, info os.FileInfo, err error) error {
			forkIfNeeded(path, info, &wg, partialResults)
			return nil
		})

	finalResult := joinResults(partialResults)

	wg.Wait()

	close(partialResults)

	result := <-finalResult
	fmt.Printf("%s has the deepest nested code block of %d\n",
		result.file, result.level)
}

func deepestNestedBlock(filename string) CodeDepth {
	code, _ := os.ReadFile(filename)
	var max float64
	var level float64
	for _, c := range code {
		switch c {
		case '{':
			level += 1
			max = math.Max(max, level)
		case '}':
			level -= 1
		}
	}
	return CodeDepth{filename, int(max)}
}

func forkIfNeeded(path string, info os.FileInfo, wg *sync.WaitGroup, results chan CodeDepth) {
	if !info.IsDir() && strings.HasSuffix(path, ".go") {
		wg.Add(1)
		go func() {
			results <- deepestNestedBlock(path)
			wg.Done()
		}()
	}
}

func joinResults(partialResults chan CodeDepth) chan CodeDepth {
	finalResult := make(chan CodeDepth)
	max := CodeDepth{"", 0}
	go func() {
		for pr := range partialResults {
			if pr.level > max.level {
				max = pr
			}
		}
		finalResult <- max
	}()
	return finalResult
}
