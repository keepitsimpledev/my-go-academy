package print10_test

import (
	"bytes"
	"go_academy/to-do-app/1.10"
	"os"
	"testing"
)

func TestPrintThings(t *testing.T) {
	want := `1. amp
2. banana
3. car
4. doorknob
5. egg
6. feather
7. gold
8. hanger
9. inchworm
10. jellybean
`

	buffer := bytes.Buffer{}
	print10.PrintThings(&buffer,
		"amp", "banana", "car", "doorknob", "egg", "feather", "gold", "hanger", "inchworm", "jellybean")

	got := buffer.String()

	if got != want {
		t.Errorf("got:\n%s\n\nwant:\n%s", got, want)
	}
}

func ExamplePrintThings() {
	print10.PrintThings(os.Stdout, "cat", "dog", "ant", "fox", "cow", "bee", "elk", "hen", "emu", "bat")
	// Output:
	// 1. cat
	// 2. dog
	// 3. ant
	// 4. fox
	// 5. cow
	// 6. bee
	// 7. elk
	// 8. hen
	// 9. emu
	// 10. bat
}

func BenchmarkPrintThings(b *testing.B) {
	for i := 0; i < b.N; i++ {
		buffer := bytes.Buffer{}
		print10.PrintThings(&buffer, "cat", "dog", "ant", "fox", "cow", "bee", "elk", "hen", "emu", "bat")
	}
}
