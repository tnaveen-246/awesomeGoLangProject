package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts)
	} else {
		for _, fileName := range files {
			f, err := os.Open(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}

func countLines(stdin *os.File, counts map[string]int) {
	input := bufio.NewScanner(stdin)
	fmt.Println(stdin.Name())
	for input.Scan() {
		text := input.Text()
		if text == "EOF" {
			break
		}
		counts[text]++
	}
}
