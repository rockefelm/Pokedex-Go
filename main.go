package main

import (
	"time"

	"github.com/rockefelm/Pokedex-Go/structs"
)
func main(){
	client := structs.NewClient(5 * time.Second)
	cfg := &structs.Config{
		PokeapiClient: client,
	}

	repl(cfg)
}