package smi

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{12.0, 13.0}

	got := Perimeter(rectangle)
	want := 50.0

	if got != want {
		t.Errorf("got %.2f want %.2f", got, want)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		shape        Shape
		expectedArea float64
	}{
		{Rectangle{Height: 11.0, Width: 11.0}, 121},
		{Circle{Radius: 10.0}, 314.1592653589793},
		{Triangle{Base: 12, Height: 6}, 36.0},
	}

	for _, areaTest := range areaTests {
		actualArea := areaTest.shape.Area()

		if actualArea != areaTest.expectedArea {
			t.Errorf("expected: %g. actual: %g, input: %v", areaTest.expectedArea, actualArea, areaTest.shape)
		}
	}
}
