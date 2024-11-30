package race

import (
	"fmt"
	"io"
	"sync"
)

const count = 1000
const parity = 2

type AlertingValue struct {
	writer            io.Writer
	value, numUpdates int
}

func (a *AlertingValue) updateAndAlert(addAmount int) {
	a.value += addAmount
	a.numUpdates++
	fmt.Fprintf(a.writer, "value: %d. numUpdates: %d\n", a.value, a.numUpdates)
}

func UpdateNumber(writer io.Writer) {
	alertingValue := AlertingValue{writer: writer, value: 0, numUpdates: 0}

	var waitGroup sync.WaitGroup

	go func(av *AlertingValue, wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%parity == 0 {
				av.updateAndAlert(parity)
			}
		}
		wg.Done()
	}(&alertingValue, &waitGroup)

	go func(av *AlertingValue, wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%parity == 1 {
				av.updateAndAlert(i % parity)
			}
		}
		wg.Done()
	}(&alertingValue, &waitGroup)

	waitGroup.Add(parity)
	waitGroup.Wait()
}
