package print10

import (
	"fmt"
	"io"
)

func PrintThings(writer io.Writer, things ...string) {
	for i := 0; i < 10; i++ {
		fmt.Fprintf(writer, "%d. %s\n", i+1, things[i])
	}
}
