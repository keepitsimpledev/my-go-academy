package sync

import (
	"fmt"
	"io"
	"sync"
)

type AlertingValue struct {
	mu                sync.Mutex
	writer            io.Writer
	value, numUpdates int
}

func (a *AlertingValue) updateAndAlert(addAmount int) {
	a.mu.Lock()
	a.value += addAmount
	a.numUpdates++
	fmt.Fprintf(a.writer, "value: %d. numUpdates: %d\n", a.value, a.numUpdates)
	a.mu.Unlock()
}

func (a *AlertingValue) GetNumUpdates() int {
	a.mu.Lock()
	numUpdates := a.numUpdates
	a.mu.Unlock()

	return numUpdates
}

func (a *AlertingValue) GetValue() int {
	a.mu.Lock()
	value := a.value
	a.mu.Unlock()

	return value
}
