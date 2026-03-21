package main

import (
	"fmt"
	"bufio"
	"os"
)

type cliCommand struct {
	name 					string
	description		string
	callback			func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(commands map[string]cliCommand) func() error {
	return func() error {
		fmt.Println("Usage:")
		for i := range commands {
			fmt.Println(commands[i].name, ":", commands[i].description)
		}
		return nil
	}
}

func main() {
	// Create Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Create valid commands
	commands := map[string]cliCommand {
		"exit" : {
			name: 				"exit",
			description: 	"Exit the Pokedex",
			callback:			commandExit,
		},
		"help" : {
			name:					"help",
			description:	"Displays a help message",
			callback:			nil,
		},
	}

	// Assign commandHelp(commands) to help after commands is created
	helpCmd := commands["help"]
	helpCmd.callback = commandHelp(commands)
	commands["help"] = helpCmd
	
	// Welcome message
	fmt.Println("Welcome to the Pokedex!")
	// Main Loop
	for true {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		cInput := cleanInput(input)
		_, exists := commands[cInput[0]]
		if exists == false {
			fmt.Println("Unknown command")
		} else {
			if err := commands[cInput[0]].callback(); err != nil {
				fmt.Println(err)
			}
		}
	}
}

