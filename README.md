# Task Tracker in GO

A simple task tracker CLI program in Go that only uses standard libraries.

https://roadmap.sh/projects/task-tracker



## Installation 

### What you need:

- go 1.24.x

1. To install, clone the repo.

`git clone $LINK`

2. In the folder where you downloaded the repo, run:

`go build ./...`

3. Then run it.

`go run ./...`


## First time running the program:

When running the task tracker for the first time, if there is no `todo.json` present within the same directory as the program, it will automatically create a `.json` file.

```
JSON file does not exist! Creating a new one...
task-cli:
```



## Available commmands

```
add {description} -- Adds a new task with a given description.
update {Id} {description} -- updates task with a new description.
delete {Id} -- deletes task.
list {status} -- lists all tasks based on status (todo, in-progress, done). if no argument is provided, lists all tasks by default.
mark-in-progress {Id} -- marks a task as in progress.
mark-done {Id} -- marks a task as done.
help -- lists the available commands.
exit -- exits the program, automatically saves changes to todos.json
```