package main

import "strings"

func main(){
	println("Hello, World!")
}

func cleanInput(text string) []string {
	var cleaned []string
	for _, line := range strings.Split(text, " ") {
		line = strings.TrimSpace(line)
		if line != "" {
			cleaned = append(cleaned, line)
		}
	}
	return cleaned
}