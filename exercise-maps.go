package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	result := make(map[string]int)
	slices := strings.Fields(s)
	for _, slice := range slices {
		result[slice]++
	}
	return result
}

func main() {
	wc.Test(WordCount)
}
