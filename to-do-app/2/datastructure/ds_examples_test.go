package ds

import "fmt"

func ExampleTodoList_Add() {
	var list TodoList

	list.Add("take a nap", "not started")

	fmt.Print(list.tasks[0])
	// Output: take a nap - not started
}

func ExampleTodoList_Get() {
	var list TodoList
	list.tasks = append(list.tasks, task{"wash dishes", "complete"})

	entry, _ := list.Get(0)

	fmt.Print(entry)
	// Output: wash dishes - complete
}

func ExampleTodoList_GetAll() {
	var list TodoList
	list.tasks = append(list.tasks,
		task{"wash dishes", "complete"},
		task{"submit taxes", "not started"},
		task{"take a nap", "in-progress"},
	)

	output := list.GetAll()

	fmt.Print(output)
	// Output: To-Do list:
	// 1. wash dishes - complete
	// 2. submit taxes - not started
	// 3. take a nap - in-progress
}

func ExampleTodoList_Update() {
	var list TodoList
	list.tasks = append(list.tasks, task{"take a nap", "not started"})

	err := list.Update(0, "take a nap", "complete")
	if err != nil {
		panic(err)
	}

	fmt.Print(list.tasks[0].String())
	// Output: take a nap - complete
}

func ExampleTodoList_Delete() {
	var list TodoList
	list.tasks = append(list.tasks, task{"wash dishes", "complete"})

	err := list.Delete(0)
	if err != nil {
		panic(err)
	}

	fmt.Print(list.GetAll())
	// Output: To-Do list is empty
}
