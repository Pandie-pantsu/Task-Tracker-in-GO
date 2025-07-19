package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func startRepl(todos *Todos) {

	scanner := bufio.NewScanner(os.Stdin)

	for {

		fmt.Print("task-cli: ")
		scanner.Scan()
		text := scanner.Text()

		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args1 := cleaned[1:]
		args2 := []string{}
		if len(cleaned) > 1 {
			args2 = cleaned[2:]
		}

		if commandName == "exit" {
			break
		}

		switch commandName {
		case "help":
			help()
		case "add":
			if len(args1) != 0 {
				todos.createTask(strings.Join(args1, " "))
			} else {
				fmt.Println("No task description given.")
			}
		case "update":
			if len(args1) != 0 {
				idx, _ := strconv.Atoi(args1[0])
				if idx == 0 {
					fmt.Println("ID 0 is not valid.")
				} else {
					todos.update(idx-1, strings.Join(args2, " "))
				}
			} else {
				fmt.Println("Missing Id to update.")
			}

		case "delete":
			if len(args1) != 0 {

				idx, _ := strconv.Atoi(args1[0])

				if idx <= 0 {
					fmt.Println("ID 0 is not valid.")
				} else {
					todos.delete(idx - 1)
				}

			} else {
				fmt.Println("Missing Id to delete.")
			}
		case "mark-in-progress":
			if len(args1) != 0 {

				idx, _ := strconv.Atoi(args1[0])

				if idx <= 0 {
					fmt.Println("Invalid ID.")
				} else {
					todos.inProgress(idx - 1)
				}

			} else {
				fmt.Println("Missing Id to mark as in-progress.")
			}
		case "mark-done":
			if len(args1) != 0 {

				idx, _ := strconv.Atoi(args1[0])

				if idx <= 0 {
					fmt.Println("ID 0 is not valid.")
				} else {
					todos.done(idx - 1)
				}

			} else {
				fmt.Println("Missing Id to mark as done.")
			}
		case "list":
			if len(args1) != 0 {
				switch args1[0] {
				case "done":
					todos.list_done()
				case "in-progress":
					todos.list_progress()
				case "todo":
					todos.list_todo()
				default:
					fmt.Println("not a valid status argument. must be 'done', 'in-progress', or 'todo'.")
				}
			} else {
				todos.list_tasks()
			}
		default:
			fmt.Println("hey man thats not a real command.")
		}
	}
}

type cliCommand struct {
	name        string
	description string
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Prints the help menu",
		},
		"add": {
			name:        "add {description}",
			description: "adds a new task, autosaves into todos.json",
		},
		"update": {
			name:        "update {id} {descrption}",
			description: "updates an existing task with a new description",
		},
		"delete": {
			name:        "delete {id}",
			description: "deletes a task from a given ID",
		},
		"mark-in-progress": {
			name:        "mark-in-progress {id}",
			description: "marks a task as in progress",
		},
		"mark-done": {
			name:        "mark-done {id}",
			description: "marks a task as done",
		},
		"list": {
			name:        "list {status}",
			description: "lists tasks, can list by status",
		},
		"exit": {
			name:        "exit",
			description: "exits the program, autosaves your tasks",
		},
	}
}

func cleanInput(str string) []string {

	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
