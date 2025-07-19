package main

import (
	"fmt"
	"os"
)

func main() {

	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	ok := storage.Load(&todos)
	if storage.Load(&todos) != ok {
		fmt.Print("JSON file does not exist! Creating a new one...\n")
		storage.Save(todos)
	}

	startRepl(&todos)

	fmt.Println("Exiting. . .")
	storage.Save(todos)
	os.Exit(0)

}
