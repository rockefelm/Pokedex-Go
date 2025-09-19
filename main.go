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
		command := cleanInput(scanner.Text())

		if len(command) == 0 {
			continue
		}
		if cmd, exists := commandRegistry[command[0]]; !exists {
			fmt.Printf("Unknown command: %v\n", command[0])
		} else {
			err := cmd.callback()
			if err != nil {
				fmt.Printf("Error executing command %v: %v\n", command[0], err)
			}
		}
	}
}