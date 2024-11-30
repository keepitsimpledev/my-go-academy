package race

import (
	"fmt"
	"io"
	"sync"
)

const count = 1000
const parity = 2

type AlertingValue struct {
	writer io.Writer
	value  int
}

func (a AlertingValue) updateAndAlert(newValue int) {
	a.value = newValue
	fmt.Fprintln(a.writer, a.value)
}

func UpdateNumber(writer io.Writer) {
	alertingValue := AlertingValue{writer: writer, value: 0}

	var waitGroup sync.WaitGroup

	go func(av *AlertingValue, wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%2 == 0 {
				av.updateAndAlert(i)
			}
		}
		wg.Done()
	}(&alertingValue, &waitGroup)

	go func(av *AlertingValue, wg *sync.WaitGroup) {
		for i := 0; i < count; i++ {
			if i%2 == 1 {
				av.updateAndAlert(i)
			}
		}
		wg.Done()
	}(&alertingValue, &waitGroup)

	waitGroup.Add(parity)
	waitGroup.Wait()
}
