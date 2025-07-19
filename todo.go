package main

import (
	"errors"
	"fmt"
	"time"
)

// Properties of the task object
type Todo struct {
	Id          int
	Description string
	Status      string
	CreatedAt   string
	UpdatedAt   string
}

type Todos []Todo // Creates a Todos slice. Delicious!

func (todos *Todos) createTask(args string) { // this function adds a new task

	time := timeFormatter(time.Now())

	todo := Todo{
		Id:          len(*todos) + 1,
		Description: args,
		Status:      "todo",
		CreatedAt:   time,
		UpdatedAt:   "",
	}

	*todos = append(*todos, todo)

	fmt.Printf("Task added successfully (ID: %+v)\n", todo.Id)

}

func timeFormatter(time time.Time) string {
	formattedTime := time.Format("01-02-2006 @ 03:04:05 PM")
	return formattedTime
}

// checks to see if index is valid, otherwise throws an error
func (todos *Todos) verifyIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}

	return nil
}

// update function
func (todos *Todos) update(index int, description string) error {

	t := *todos

	if err := t.verifyIndex(index); err != nil {
		return err
	}

	t[index].Description = description

	updatedTime := timeFormatter(time.Now())
	t[index].UpdatedAt = updatedTime

	fmt.Printf("Updated Task ID: %+v. \n", index+1)

	return nil
}

// delete function
func (todos *Todos) delete(index int) error {

	t := *todos

	if err := t.verifyIndex(index); err != nil {
		return err
	}

	*todos = append(t[:index], t[index+1:]...)
	if index <= len(*todos) {
		for idx := index; idx < len(*todos); idx++ {
			t[idx].Id--
		}
	}

	return nil
}

func (todos *Todos) inProgress(index int) error {

	t := *todos
	s := "in progress"

	if err := t.verifyIndex(index); err != nil {
		return err
	}

	updatedTime := timeFormatter(time.Now())
	t[index].UpdatedAt = updatedTime

	t[index].Status = s

	return nil
}

func (todos *Todos) done(index int) error {
	t := *todos
	s := "done"

	if err := t.verifyIndex(index); err != nil {
		return err
	}

	updatedTime := timeFormatter(time.Now())
	t[index].UpdatedAt = updatedTime

	t[index].Status = s

	return nil
}
