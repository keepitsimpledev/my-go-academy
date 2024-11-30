package ds

import "fmt"

type Task struct {
	item, status string
}

func NewTask(item, status string) Task {
	return Task{
		item:   item,
		status: status,
	}
}

func (t Task) GetItem() string {
	return t.item
}

func (t Task) GetStatus() string {
	return t.status
}

func (t Task) String() string {
	return fmt.Sprintf("%s - %s", t.item, t.status)
}
