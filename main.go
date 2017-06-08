package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func readLineByLine(target string, pathToFile string) {
	file, err := os.Open(pathToFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, target) {
			fmt.Printf("%v: %v\n", pathToFile, line)
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	root := flag.Arg(0)
	fileType := flag.Arg(1)
	searchFor := flag.Arg(2)

	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if f != nil && f.Mode().IsRegular() && strings.HasSuffix(f.Name(), fileType) {
			readLineByLine(searchFor, path)
		}
		return nil
	})
	if err != nil {
		fmt.Printf("filepath.Walk() returned %v\n", err)
	}
}
