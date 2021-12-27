package utils

import (
	"math"
	"strconv"
)

type Coordinate struct {
	x float64
	y float64
}

func NewCoord(x, y float64) Coordinate {
	return Coordinate{x: x, y: y}
}

func (c *Coordinate) X() float64 {
	return c.x
}

func (c *Coordinate) Y() float64 {
	return c.y
}

func (c Coordinate) Move(point Coordinate) Coordinate {
	return Coordinate{x: c.x + point.x, y: c.y + point.y}
}

func (c *Coordinate) Rotate(angle float64) Coordinate {
	return Coordinate{
		x: c.x*math.Cos(angle) - c.y*math.Sin(angle),
		y: c.x*math.Sin(angle) + c.y*math.Cos(angle),
	}
}

func (c *Coordinate) Scale(factor float64) Coordinate {
	return Coordinate{x: c.x * factor, y: c.y * factor}
}

func (c *Coordinate) DistanceTo(point Coordinate) float64 {
	return math.Sqrt(c.DistanceToSq(point))
}

func (c *Coordinate) DistanceToSq(point Coordinate) float64 {
	return math.Pow(c.x-point.x, 2) + math.Pow(c.y-point.y, 2)
}

func (c *Coordinate) Equal(point Coordinate) bool {
	return c.x == point.x && c.y == point.y
}

func (c *Coordinate) String() string {
	return "(" + strconv.FormatFloat(c.x, 'f', 2, 64) + ", " + strconv.FormatFloat(c.y, 'f', 2, 64) + ")"
}

func (c *Coordinate) Copy() Coordinate {
	return Coordinate{x: c.x, y: c.y}
}

func (c *Coordinate) Mirror(ref1, ref2 Coordinate) Coordinate {
	return Coordinate{x: 2*ref1.x - c.x, y: 2*ref1.y - c.y}
}
