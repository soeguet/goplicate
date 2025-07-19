package main

import (
	"crypto/sha256"
	"fmt"
	"os"
	"path/filepath"
)

type Datei struct {
	name string
	path string
	hash string
}

var datei []Datei
var counter int

func recusivlyReadFiles(currentPath string, path string) {
	entries, err := os.ReadDir(currentPath)
	if err != nil {
		return
	}
	for _, d := range entries {
		fullPath := filepath.Join(currentPath, d.Name())
		if d.IsDir() {
			recusivlyReadFiles(fullPath, path+"/"+d.Name())
		} else {
			file, err := os.ReadFile(fullPath)
			if err == nil {
				h := sha256.New()
				h.Write(file)
				bs := h.Sum(nil)

				newDatei := Datei{
					name: d.Name(),
					path: fullPath,
					hash: fmt.Sprintf("%x", bs),
				}

				datei = append(datei, newDatei)

				counter++
			}
		}
	}
}

func main() {

	recusivlyReadFiles(".", ".")

	for _, d := range datei {
		fmt.Printf("File name: %s\nFile path: %s\nHash: %s\n\n", d.name, d.path, d.hash)
	}
	fmt.Println("_____________")

	duplikate := make(map[string][]Datei)

	for _, d := range datei {
		duplikate[d.hash] = append(duplikate[d.hash], d)
	}

	for mapKey, mapValue := range duplikate {

		if len(mapValue) > 1 {

			fmt.Printf("Duplicates for Hash: %s\n", mapKey)
			for index, value := range mapValue {
				fmt.Printf("%d: %s\n", index, value.path)
			}
		}

	}

	fmt.Printf("Total number of files: %d\n", len(datei))
	fmt.Println("Counter: ", counter)
}
