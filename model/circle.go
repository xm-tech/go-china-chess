package model

import "image/color"

type Circle struct {
	X        float64
	Y        float64
	R        float64
	Color    color.Color
	X_origin float64
}

func (circle *Circle) Run(xMax float64) {
	if circle.X >= xMax {
		circle.X = circle.X_origin
	} else {
		circle.X += 1
	}
}

func NewCircle(x, y, r float64, color color.Color) *Circle {
	circle := &Circle{
		X:     x,
		Y:     y,
		R:     r,
		Color: color,
	}
	circle.X_origin = circle.X
	return circle
}
