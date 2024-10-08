package clockface

import (
	"math"
	"time"
)

/// TODO: finish chapter

type Point struct {
	X float64
	Y float64
}

func SecondHand(t time.Time) Point {
	return Point{150, 60}
}

func secondsInRadians(t time.Time) float64 {
	return (math.Pi / (30 / (float64(t.Second()))))
}

func secondHandPoint(t time.Time) Point {
	return Point{0, -1}
}
