package ds

import "testing"

const item = "wash dishes"
const status = "not started"

func TestTask(t *testing.T) {
	task := Task{item, status}

	assert(t, task.GetItem(), item)
	assert(t, task.GetStatus(), status)
	assert(t, task.String(), "wash dishes - not started")
}

func TestNewList(t *testing.T) {
	todoList := NewList()

	t.Run("test list length", func(t *testing.T) {
		want := 0
		got := len(todoList.tasks)
		assert(t, got, want)
	})
}

func TestAdd(t *testing.T) {
	var todoList TodoList

	todoList.Add(item, status)

	t.Run("test length", func(t *testing.T) {
		want := 1
		got := len(todoList.tasks)
		assert(t, got, want)
	})

	t.Run("test item", func(t *testing.T) {
		want := item
		got := todoList.tasks[0].item
		assert(t, got, want)
	})

	t.Run("test status", func(t *testing.T) {
		want := status
		got := todoList.tasks[0].status
		assert(t, got, want)
	})
}

func TestGet(t *testing.T) {
	var todoList TodoList

	testTask := Task{item, status}
	todoList.tasks = append(todoList.tasks, testTask)

	t.Run("standard get", func(t *testing.T) {
		got, _ := todoList.Get(0)
		assert(t, got, testTask)
	})

	t.Run("out-of-bounds", func(t *testing.T) {
		_, err := todoList.Get(1)
		assert(t, err.Error(), "index out-of-bounds")
	})
}

func TestGetAll(t *testing.T) {
	t.Run("empty list", func(t *testing.T) {
		// arrange
		var todoList TodoList

		todoList.tasks = append(todoList.tasks, Task{"item 1", "status 1"}, Task{"item 2", "status 2"})
		want := "To-Do list:\n1. item 1 - status 1\n2. item 2 - status 2\n"

		// act
		got := todoList.GetAll()

		// assert
		if got != want {
			t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		// arrange
		var todoList TodoList

		// act
		got := todoList.GetAll()

		// assert
		want := "To-Do list is empty\n"
		if got != want {
			t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
		}
	})
}

func TestUpdate(t *testing.T) {
	var todoList TodoList

	t.Run("out-of-bounds", func(t *testing.T) {
		err := todoList.Update(-1, "", "")
		assert(t, err.Error(), "index out-of-bounds")
	})

	t.Run("out-of-bounds", func(t *testing.T) {
		err := todoList.Update(0, "", "")
		assert(t, err.Error(), "index out-of-bounds")
	})

	t.Run("update", func(t *testing.T) {
		todoList.tasks = append(todoList.tasks, Task{item, status})

		err := todoList.Update(0, "do laundry", "in-progress")

		assert(t, err, nil)

		task := todoList.tasks[0]

		assert(t, task.item, "do laundry")
		assert(t, task.status, "in-progress")
	})
}

func TestDelete(t *testing.T) {
	t.Run("out-of-bounds", func(t *testing.T) {
		var todoList TodoList

		err := todoList.Delete(0)

		assert(t, err.Error(), "index out-of-bounds")
	})

	t.Run("delete one of one", func(t *testing.T) {
		var todoList TodoList

		todoList.tasks = append(todoList.tasks, Task{item, status})

		err := todoList.Delete(0)

		assert(t, err, nil)
		assert(t, len(todoList.tasks), 0)
	})

	task1 := Task{"wash dishes", "not started"}
	task2 := Task{"do laundry", "in-progress"}
	task3 := Task{"take a nap", "complete"}

	populate3Tasks := func() *TodoList {
		var todoList TodoList

		todoList.tasks = append(todoList.tasks, task1, task2, task3)

		return &todoList
	}

	t.Run("delete first", func(t *testing.T) {
		todoList := populate3Tasks()

		err := todoList.Delete(0)

		assert(t, err, nil)
		assertTaskList(t, todoList.tasks, []Task{task2, task3})
	})

	t.Run("delete middle", func(t *testing.T) {
		todoList := populate3Tasks()

		err := todoList.Delete(1)

		assert(t, err, nil)
		assertTaskList(t, todoList.tasks, []Task{task1, task3})
	})

	t.Run("delete last", func(t *testing.T) {
		todoList := populate3Tasks()

		err := todoList.Delete(2)

		assert(t, err, nil)
		assertTaskList(t, todoList.tasks, []Task{task1, task2})
	})
}

func assert(tb testing.TB, got, want any) {
	tb.Helper()

	if got != want {
		tb.Errorf("got: %v. want: %v.", got, want)
	}
}

func assertTaskList(tb testing.TB, got, want []Task) {
	tb.Helper()

	if len(got) != len(want) {
		tb.Errorf("got length: %v. want length: %v.", len(got), len(want))
	} else {
		for i := 0; i < len(got); i++ {
			if got[i] != want[i] {
				tb.Errorf("got: '%v'. want: '%v'.", got[i].String(), want[i].String())
			}
		}
	}
}
