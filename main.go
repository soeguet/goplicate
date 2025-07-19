package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"
)

type Files struct {
	name string
	path string
	hash string
}

func recursivelyReadFiles(currentPath string, wg *sync.WaitGroup, ch chan<- Files) {
	defer wg.Done()

	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return
	}

	for _, d := range entries {
		fullPath := filepath.Join(currentPath, d.Name())

		if d.IsDir() {
			wg.Add(1)
			go recursivelyReadFiles(fullPath, wg, ch)
		} else {
			file, err := os.ReadFile(fullPath)
			if err == nil {
				h := sha256.New()
				h.Write(file)
				hash := fmt.Sprintf("%x", h.Sum(nil))

				ch <- Files{
					name: d.Name(),
					path: fullPath,
					hash: hash,
				}
			}
		}
	}
}

func main() {
	now := time.Now()

	files, counter := handleGoroutines()

	printOverview(files)
	determineDuplicates(files)
	printResult(files, counter, now)
}

func handleGoroutines() ([]Files, int) {
	filesChan := make(chan Files)
	var wg sync.WaitGroup
	var files []Files
	var counter int

	go func() {
		for d := range filesChan {
			files = append(files, d)
			counter++
		}
	}()

	wg.Add(1)
	go recursivelyReadFiles(".", &wg, filesChan)

	wg.Wait()
	close(filesChan)
	return files, counter
}

func printResult(files []Files, counter int, now time.Time) {
	fmt.Printf("Total number of files: %d\n", len(files))
	fmt.Println("Counter: ", counter)
	fmt.Printf("Time passed: %v\n", time.Since(now))
}

func printOverview(files []Files) {
	for _, d := range files {
		fmt.Printf("File name: %s\nFile path: %s\nHash: %s\n\n", d.name, d.path, d.hash)
	}
	fmt.Println("_____________")
}

func determineDuplicates(files []Files) {
	duplicates := make(map[string][]Files)
	for _, d := range files {
		duplicates[d.hash] = append(duplicates[d.hash], d)
	}

	for hash, files := range duplicates {
		if len(files) > 1 {
			fmt.Printf("Duplicates for Hash: %s\n", hash)
			for i, f := range files {
				fmt.Printf("%d: %s\n", i, f.path)
			}
		}
	}
}
