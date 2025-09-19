package main

import "fmt"

func commandHelp() error {
	for key, value := range commandRegistry {
		fmt.Print("Welcome to the Pokedex!\nUsage:\n\n\n")
		fmt.Printf("%v: %v", key, value.description)
	}
	return nil
}