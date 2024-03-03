# Task Manager CLI

This is a simple CLI application to manage tasks a.k.a simple Todo CLI application. Users can add tasks, list all tasks, mark tasks as complete, and remove tasks from the list.

## Build from Source

To build the application from source, clone the repository and build the binary using Go:

```bash
gh repo clone Kei-K23/go-todo-cli
cd go-todo-cli
go build main.go
```

## Usage

Build `main.go` file with `go build main.go` command to get executable program.

### Add a new task

Use the -add flag followed by the task description to add a new task.

```bash
    ./main -add "To learn Go programming"
```

### List All Tasks

To list all tasks, use the -list flag.

```bash
    ./main -list
```

### Mark a Task as Complete

Specify the task ID using the -complete flag to mark it as complete.

```bash
 ./main -complete 1
```

### Remove a Task

Use the -remove flag followed by the task ID to remove it from the list.

```bash
 ./main -remove 1
```
