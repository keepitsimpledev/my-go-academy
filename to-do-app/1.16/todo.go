package todoprint

type Todo struct {
	task, status string
}

func NewTodo(task, status string) Todo {
	return Todo{task: task, status: status}
}

func (t Todo) GetTask() string {
	return t.task
}

func (t Todo) GetStatus() string {
	return t.status
}
