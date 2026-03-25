package main

import (
	"fmt"
	"errors"
)

func commandExplore(cfg *config, areaName string) error {
	if areaName == "" {
		return errors.New("you must provide a location name")
	}
	location, err := cfg.pokeapiClient.GetLocation(areaName)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", location.Name)
	fmt.Println("Found Pokemon: ")
	for _, encounter := range location.PokemonEncounters {
		fmt.Printf(" - %s\n", encounter.Pokemon.Name)
	}
	return nil
}