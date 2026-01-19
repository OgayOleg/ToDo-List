package main

import (
	"todolist/list"
	"todolist/scaner"
)

func main() {

	todolist := list.NewStorageTask()
	Task := scaner.NewScanner(todolist)
	Task.Start()
}
