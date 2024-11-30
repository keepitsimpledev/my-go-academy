package integers

import "fmt"
import "testing"

func TestAdder(t *testing.T) {
	assertSum(t, 2, 2, 4)
	assertSum(t, 2, 3, 5)
	assertSum(t, 3, 3, 6)
	assertSum(t, 2, 4, 6)
	assertSum(t, 6, 45, 51)
}

func assertSum(tb testing.TB, addendA, addendB int, sum int) {
	tb.Helper()

	actual := Add(addendA, addendB)
	expected := sum

	if actual != expected {
		tb.Errorf("expected '%d' but got '%d'", expected, actual)
	}
}

func ExampleAdd() {
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6
}
