package sync_test

import (
	"bytes"
	"go_academy/to-do-app/1.15"
	"testing"
)

func TestUpdateNumber(t *testing.T) {
	buffer := bytes.Buffer{}
	alertingValue := sync.UpdateNumber(&buffer)

	if len(buffer.String()) == 0 {
		t.Error("nothing written")
	}

	wantUpdates := 1000
	gotUpdates := alertingValue.GetNumUpdates()

	if wantUpdates != gotUpdates {
		t.Errorf("got: %d. want %d.", gotUpdates, wantUpdates)
	}
}

func BenchmarkUpdateNumber(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		sync.UpdateNumber(&buffer)
	}
}
