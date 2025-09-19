package main

import (
	"time"

	"github.com/rockefelm/Pokedex-Go/structs"
)
func main(){
	client := structs.NewClient(5 * time.Second, time.Minute*5)
	cfg := &structs.Config{
		PokeapiClient: client,
	}

	repl(cfg)
}