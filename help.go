package main

import "fmt"

func help() {
	fmt.Println("Help Menu")
	availableCommands := getCommands()

	for _, cmd := range availableCommands {
		fmt.Printf(" - %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("")
}
