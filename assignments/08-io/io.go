package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const Filename = "cities.txt"

func CreateFile() *os.File {
	file, err := os.Create(Filename)
	if err != nil {
		panic(err)
	}

	return file
}

func WriteToFile(content string, file *os.File) {
	_, err := file.Write([]byte(content))
	if err != nil {
		DeleteFile()
		panic(err)
	}
}

func ReadFromFile() string {
	content, err := os.ReadFile(Filename)
	if err != nil {
		DeleteFile()
		panic(err)
	}

	return string(content)
}

func DeleteFile() {
	os.Remove(Filename)
}

func WriteAndReadFromFile(s []string) []string {
	file := CreateFile()
	combined := strings.Join(s, "\n")

	WriteToFile(combined, file)

	fileContents := ReadFromFile()
	fileContentsList := strings.Split(fileContents, "\n")

	sort.Strings(fileContentsList)
	DeleteFile()

	return fileContentsList
}

func main() {
	citiesList := []string{"Abu Dhabi", "London", "Washington D.C.", "Montevideo", "Vatican City", "Caracas", "Hanoi"}
	sortedCities := WriteAndReadFromFile(citiesList)

	for i := 0; i < len(sortedCities); i++ {
		fmt.Printf("%d) %s\n", i+1, sortedCities[i])
	}
}
