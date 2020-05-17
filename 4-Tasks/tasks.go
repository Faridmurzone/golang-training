package main

import "fmt"

type task struct {
	name        string
	description string
	completed   bool
}

func (t task) handleCompleted() {
	t.completed = true
}

func (t task) handleDescription(description string) {
	t.description = description
}

func (t task) handleName(name string) {
	t.name = name
}

func main() {
	t := task{
		name:        "Finish this GO exercise",
		description: "I have to make a struct and receiver function for create tasks in a list",
		completed:   false,
	}
	fmt.Printf("%+v\n", t)
	t.handleCompleted()
	t.handleDescription("Created a tasks update with structs and receivers")
	t.handleName("Finished")
	fmt.Printf("%+v\n", t)
}
