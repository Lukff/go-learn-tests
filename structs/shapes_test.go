package structs

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{10, 10}, hasPerimeter: 40},
		{name: "Circle", shape: Circle{5}, hasPerimeter: 31.41592653589793},
		{name: "Triangle", shape: Triangle{A: 3, B: 4, C: 5}, hasPerimeter: 12.0},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %.2f want %.2f", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Width: 12, Heigth: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{A: 3, B: 4, C: 5}, hasArea: 6.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g want %g", tt.shape, got, tt.hasArea)
			}
		})
	}

}
