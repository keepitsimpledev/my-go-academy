package ds

import (
	"fmt"
	"testing"
)

func BenchmarkAdd(b *testing.B) {
	var todoList TodoList

	for i := 0; i < b.N; i++ {
		item := fmt.Sprintf("item: %d", i)
		status := fmt.Sprintf("status: %d", i)
		todoList.Add(item, status)
	}
}

func BenchmarkAddAndGet(b *testing.B) {
	var todoList TodoList

	for i := 0; i < b.N; i++ {
		item := fmt.Sprintf("item: %d", i)
		status := fmt.Sprintf("status: %d", i)
		todoList.Add(item, status)
	}

	for i := 0; i < b.N; i++ {
		_, err := todoList.Get(i)
		panicIfErr(err)
	}
}

func BenchmarkAddAndGetAll(b *testing.B) {
	var todoList TodoList

	for i := 0; i < b.N; i++ {
		item := fmt.Sprintf("item: %d", i)
		status := fmt.Sprintf("status: %d", i)
		todoList.Add(item, status)
	}

	todoList.GetAll()
}

func BenchmarkAddAndUpdate(b *testing.B) {
	var todoList TodoList

	for i := 0; i < b.N; i++ {
		item := fmt.Sprintf("item: %d", i)
		status := fmt.Sprintf("status: %d", i)
		todoList.Add(item, status)
	}

	for i := 0; i < b.N; i++ {
		item := fmt.Sprintf("item updated: %d", i)
		status := fmt.Sprintf("status updated: %d", i)
		err := todoList.Update(i, item, status)
		panicIfErr(err)
	}
}

// commented because this runs for a long time - is it indicative of an issue?
// func BenchmarkAddAndDelete(b *testing.B) {
// 	var todoList TodoList

// 	for i := 0; i < b.N; i++ {
// 		item := fmt.Sprintf("item: %d", i)
// 		status := fmt.Sprintf("status: %d", i)
// 		todoList.Add(item, status)
// 	}

// 	for i := 0; i < b.N; i++ {
// 		err := todoList.Delete(0)
// 		panicIfErr(err)
// 	}
// }

func panicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}
