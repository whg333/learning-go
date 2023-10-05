package main

import (
	"fmt"
	"learning-go/whg/input/files"
)

func main() {
	lines := files.GetLines("votes.txt")
	nameCounts := make(map[string]int)
	for _, line := range lines {
		nameCounts[line]++
	}
	for name, count := range nameCounts {
		fmt.Printf("Votes for %s: %d\n", name, count)
	}
}
