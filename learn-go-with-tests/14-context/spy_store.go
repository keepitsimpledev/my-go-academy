package context

import (
	"context"
	"fmt"
	"log"
	"testing"
	"time"
)

const numberOfMillisToWait = 10

type SpyStore struct {
	response string
	t        *testing.T
}

func (s *SpyStore) Fetch(ctx context.Context) (string, error) {
	data := make(chan string, 1)

	go func() {
		var result string

		for _, c := range s.response {
			select {
			case <-ctx.Done():
				log.Println("spy store got cancelled")
				return
			default:
				time.Sleep(numberOfMillisToWait * time.Millisecond)

				result += string(c)
			}
		}
		data <- result
	}()

	select {
	case <-ctx.Done():
		return "", fmt.Errorf("%v", ctx.Err())
	case res := <-data:
		return res, nil
	}
}
