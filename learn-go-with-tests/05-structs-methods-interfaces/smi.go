package smi

// number of times the height and width of a rectangle contributes towards its perimeter
const perimeterFactor = 2

type Shape interface {
	Area() float64
}

func Perimeter(rectangle Rectangle) float64 {
	return perimeterFactor * (rectangle.Height + rectangle.Width)
}

func Area(rectangle Rectangle) float64 {
	return rectangle.Height * rectangle.Width
}
