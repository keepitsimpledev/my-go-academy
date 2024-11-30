package main

import (
	"fmt"
	readandprint "go_academy/to-do-app/1.13"
	"os"
)

func main() {
	out, err := readandprint.ReadJSONTaskFiles(os.DirFS("../tasks"))
	if err != nil {
		panic(err)
	}

	fmt.Println(out)
}
