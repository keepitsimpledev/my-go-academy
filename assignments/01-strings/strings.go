package main

import "fmt"

func CombineStrings() string {
	multiple := "multiple"
	strings := "strings"
	variable := "variable"

	return fmt.Sprintf("%s %s %s", multiple, strings, variable)
}

func main() {
	fmt.Println(CombineStrings())
}
