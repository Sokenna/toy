package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, "sdtin")
	} else {
		for _, file := range files {
			data, err := os.Open(file)
			if err != nil {
				fmt.Fprintf(os.Stderr, "test1.4: %v\n", err)
				continue
			}
			countLines(data, counts, file)
		}
	}

	for i, m1 := range counts {
		for s, _ := range m1 {
			if m1[s] {
				fmt.Printf("%d\t%s\n", i, s)
			}
		}
	}
}

func countLines(f *os.File, count map[string]map[string]bool, filename string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if count[line] == nil {
			count[line] = make(map[string]bool)
		}
		count[line][filename] = true
	}
}
