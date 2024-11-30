package ds

import "fmt"

type task struct {
	item, status string
}

func (t task) getItem() string {
	return t.item
}

func (t task) getStatus() string {
	return t.status
}

func (t task) String() string {
	return fmt.Sprintf("%s - %s", t.item, t.status)
}
