package main

import (
	"fmt"
	"io"
	"os"
)

type Fullname struct {
	first, middle, last string
}

func (f Fullname) full() string {
	return fmt.Sprintf("%s %s %s", f.first, f.middle, f.last)
}

func PromptForName(in io.Reader) Fullname {
	var name Fullname

	fmt.Print("Enter first-name: ")
	fmt.Fscanln(in, &name.first)

	fmt.Print("Enter middle-name: ")
	fmt.Fscanln(in, &name.middle)

	fmt.Print("Enter surname: ")
	fmt.Fscanln(in, &name.last)

	return name
}

func FormatName(name Fullname) string {
	return fmt.Sprintf("full-name : %s\nmiddle-name : %s\nsurname : %s", name.full(), name.middle, name.last)
}

func main() {
	name := PromptForName(os.Stdin)
	fmt.Println(FormatName(name))
}
