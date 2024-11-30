//go:build !race

package race_test

import (
	"bytes"
	"go_academy/to-do-app/1.14"
	"testing"
)

func TestUpdateNumber(t *testing.T) {
	buffer := bytes.Buffer{}
	race.UpdateNumber(&buffer)

	if len(buffer.String()) == 0 {
		t.Error("nothing written")
	}
}
