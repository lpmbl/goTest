package main

import (
	"strings"
	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	var m map[string]int
	m = make(map[string]int)
	chars := strings.Fields(s)
	for _, char := range chars{
		if m[char] >= 1{
			m[char] = m[char]+1
		} else {
			m[char] = 1
		}
	}
	return m
}

func main() {
	wc.Test(WordCount)
}