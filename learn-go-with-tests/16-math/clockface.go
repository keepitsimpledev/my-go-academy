package clockface

import (
	"math"
	"time"
)

const radiansInACircle = math.Pi * float64(2)
const secondsInAMinute = 60
const minutesInAnHour = secondsInAMinute
const hoursOnAClock = 12

// A Point represents a two-dimensional Cartesian coordinate
type Point struct {
	X float64
	Y float64
}

// SecondHand is the unit vector of the second hand of an analogue clock at time `t`
// represented as a Point.
func SecondHand(t time.Time) Point {
	p := secondHandPoint(t)
	p = Point{p.X * 90, p.Y * 90}   // scale
	p = Point{p.X, -p.Y}            // flip
	p = Point{p.X + 150, p.Y + 150} // translate

	return p
}

func secondsInRadians(t time.Time) float64 {
	// textbook answer:
	// return (math.Pi / (radiansSecondsFactor / (float64(t.Second()))))
	return float64(t.Second()) * radiansInACircle / secondsInAMinute
}

func secondHandPoint(t time.Time) Point {
	return angleToPoint(secondsInRadians(t))
}

func minutesInRadians(t time.Time) float64 {
	// textbook answer:
	// return (secondsInRadians(t) / 60) +
	//     (math.Pi / (30 / float64(t.Minute())))
	return ((float64(t.Minute()) * radiansInACircle / minutesInAnHour) +
		(float64(t.Second()) * radiansInACircle / secondsInAMinute / minutesInAnHour))
}

func minuteHandPoint(t time.Time) Point {
	return angleToPoint(minutesInRadians(t) + (secondsInRadians(t) / secondsInAMinute))
}

func hoursInRadians(t time.Time) float64 {
	return (float64(t.Hour()%hoursOnAClock) * radiansInACircle / hoursOnAClock) +
		(float64(t.Minute()) * radiansInACircle / minutesInAnHour / hoursOnAClock) +
		(float64(t.Second()) * radiansInACircle / secondsInAMinute / minutesInAnHour / hoursOnAClock)
}

func hourHandPoint(t time.Time) Point {
	return angleToPoint(hoursInRadians(t))
}

func angleToPoint(angle float64) Point {
	x := math.Sin(angle)
	y := math.Cos(angle)

	return Point{x, y}
}
