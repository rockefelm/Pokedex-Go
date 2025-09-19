package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		if !scanner.Scan() {
			break
		}
		text := cleanInput(scanner.Text())

		if len(text) == 0 {
			continue
		}
		fmt.Printf("Your command was: %v\n", text[0])
	}
}