package main

import "fmt"

// returns list of things remaining
func (todos *Todos) list_tasks() error {
	for idx, t := range *todos {
		fmt.Printf("| ID# %d | Task: %s | Status: %s | Created: %s | Updated: %s |\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		idx++
	}
	return nil
}

func (todos *Todos) list_done() error {
	for idx, t := range *todos {
		if t.Status == "done" {
			fmt.Printf("| ID# %d | Task: %s | Status: %s | Created: %s | Updated: %s |\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
		idx++
	}
	return nil
}

func (todos *Todos) list_progress() error {
	for idx, t := range *todos {
		if t.Status == "in progress" {
			fmt.Printf("| ID# %d | Task: %s | Status: %s | Created: %s | Updated: %s |\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
		idx++
	}
	return nil
}

func (todos *Todos) list_todo() error {
	for idx, t := range *todos {
		if t.Status == "todo" {
			fmt.Printf("| ID# %d | Task: %s | Status: %s | Created: %s | Updated: %s |\n", t.Id, t.Description, t.Status, t.CreatedAt, t.UpdatedAt)
		}
		idx++
	}
	return nil
}
