package main

import (
	"bufio"
	"fmt"
	"os"
)

// Dup1 prints the text of each line that appears more than
// once in the standard input, preceded by its count.
func main() {
	counts := make(map[string]int)
	input := bufio.NewScanner(os.Stdin)
	for input.Scan() {
		text := input.Text()
		if text == "EOF" {
			break
		}
		counts[text]++
	}
	for key, value := range counts {
		fmt.Println(key, ":", value)
	}
}
