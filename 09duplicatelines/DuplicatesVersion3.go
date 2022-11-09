package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, fileName := range os.Args[1:] {
		data, err := os.ReadFile(fileName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "dup3: %v\n", err)
			continue
		}
		for _, val := range strings.Split(string(data), "\n") {
			counts[val]++
		}
		for id, val := range counts {
			fmt.Printf("%s\t%d\n", id, val)
		}
	}
}
