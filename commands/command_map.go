package commands

import (
	"fmt"
	"errors"

	"github.com/rockefelm/Pokedex-Go/structs"
)



func CommandMap(cfg *structs.Config) error {
	locations, err := cfg.PokeapiClient.Locations(cfg.NextLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locations.Next
	cfg.PrevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}

func CommandMapBack(cfg *structs.Config) error {
	if cfg.PrevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locations, err := cfg.PokeapiClient.Locations(cfg.PrevLocationsURL)
	if err != nil {
		return err
	}

	cfg.NextLocationsURL = locations.Next
	cfg.PrevLocationsURL = locations.Previous

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}
	return nil
}
