package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	filesNames := make(map[string][]string)
	for _, fileName := range os.Args[1:] {
		data, _ := os.ReadFile(fileName)
		for _, val := range strings.Split(string(data), "\n") {
			counts[val]++
			if !contains(filesNames[fileName], val) {
				filesNames[fileName] = append(filesNames[fileName], val)
			}
		}
	}
	for key, val := range counts {
		if val > 1 {
			fmt.Println(key, ":", val)
			for fName, strList := range filesNames {
				if contains(strList, key) {
					fmt.Println(fName)
				}
			}
		}
	}
}

func contains(list []string, val string) bool {
	for _, data := range list {
		if data == val {
			return true
		}
	}
	return false
}
