package main

import (
	"fmt"
	"errors"
)

func commandPokedex(cfg *config, args ...string) error {
	if len(cfg.caughtPokemon) == 0 {
		return errors.New("you have not caught any pokemon")
	}
	fmt.Println("Your Pokedex:")
	for name, _ := range cfg.caughtPokemon {
		fmt.Printf(" - %s\n", name)
	}

	return nil
}