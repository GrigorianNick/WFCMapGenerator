package pattern

import (
	_ "image/png"

	"github.com/grigoriannick/WFCMapGenerator/point"
)

type Orientation int

const (
	Top = Orientation(iota)
	Right
	Bottom
	Left
)

type Transformation int

const (
	RotateClockwise = Transformation(iota)
	RotateCounterClockwise
	FlipHorizontally
	FlipVertically
)

type Pattern interface {
	// Orientation is where the parameter pattern is in relation to the caller
	Accept(Pattern, Orientation) bool
	Transform(Transformation) Pattern
	GetDimensions() (int, int)
	GetEdge(Orientation) []point.Point
	SetSize(int, int)
	PointAt(int, int) point.Point
	SetPointAt(int, int, point.Point)
	Equal(Pattern) bool
	SetAcceptingEdge(Orientation, ...point.Point)
	GetAcceptingEdge(Orientation) []point.Point
}
