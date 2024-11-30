package app_test

import (
	"bytes"
	"fmt"
	app "go_academy/to-do-app/2"
	"strings"
	"testing"
)

const addInput = "1\n"
const readInput = "2\n"
const updateInput = "3\n"
const deleteInput = "4\n"
const exitInput = "5\n"

func BenchmarkAdd(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		itemInput := fmt.Sprintf("item: %d\n", i)
		statusInput := fmt.Sprintf("status: %d\n", i)
		input += fmt.Sprintf("%s%s%s", addInput, itemInput, statusInput)
	}

	input += exitInput

	outputBuffer := bytes.Buffer{}
	inputReader := strings.NewReader(input)

	app.Run(inputReader, &outputBuffer)
}

func BenchmarkAddAndRead(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		itemInput := fmt.Sprintf("item: %d\n", i)
		statusInput := fmt.Sprintf("status: %d\n", i)
		input += fmt.Sprintf("%s%s%s", addInput, itemInput, statusInput)
	}

	input += readInput
	input += exitInput

	outputBuffer := bytes.Buffer{}
	inputReader := strings.NewReader(input)

	app.Run(inputReader, &outputBuffer)
}

func BenchmarkAddAndUpdate(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		itemInput := fmt.Sprintf("item: %d\n", i)
		statusInput := fmt.Sprintf("status: %d\n", i)
		input += fmt.Sprintf("%s%s%s", addInput, itemInput, statusInput)
	}

	for i := 0; i < b.N; i++ {
		itemInput := fmt.Sprintf("item updated: %d\n", i)
		statusInput := fmt.Sprintf("status updated: %d\n", i)
		input += fmt.Sprintf("%s%s%s", updateInput, itemInput, statusInput)
	}

	input += exitInput

	outputBuffer := bytes.Buffer{}
	inputReader := strings.NewReader(input)

	app.Run(inputReader, &outputBuffer)
}

func BenchmarkAddAndDelete(b *testing.B) {
	input := ""

	for i := 0; i < b.N; i++ {
		itemInput := fmt.Sprintf("item: %d\n", i)
		statusInput := fmt.Sprintf("status: %d\n", i)
		input += fmt.Sprintf("%s%s%s", addInput, itemInput, statusInput)
	}

	for i := 0; i < b.N; i++ {
		input += fmt.Sprintf("%s1\n", deleteInput)
	}

	input += exitInput

	outputBuffer := bytes.Buffer{}
	inputReader := strings.NewReader(input)

	app.Run(inputReader, &outputBuffer)
}
