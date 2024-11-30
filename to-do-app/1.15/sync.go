package sync

import (
	"io"
	"sync"
)

const count = 1000
const parity = 2

func UpdateNumber(writer io.Writer) *AlertingValue {
	alertingValue := AlertingValue{mu: sync.Mutex{}, writer: writer, value: 0, numUpdates: 0}
	updateChannel := make(chan int)

	var waitGroup sync.WaitGroup

	waitGroup.Add(parity)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%parity == 0 {
				updateChannel <- parity
			}
		}
		wg.Done()
	}(&waitGroup)

	go func(wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%parity == 1 {
				updateChannel <- i % parity
			}
		}
		wg.Done()
	}(&waitGroup)

	for i := 0; i < count; i++ {
		updateValue := <-updateChannel
		alertingValue.updateAndAlert(updateValue)
	}

	waitGroup.Wait()

	return &alertingValue
}
