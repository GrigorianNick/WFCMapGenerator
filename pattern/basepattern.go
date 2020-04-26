package pattern

import (
	"fmt"

	"github.com/grigoriannick/WFCMapGenerator/point"
)

type BasePattern struct {
	pointMatrix  [][]point.Point
	adjacentEdge map[Orientation][]point.Point
}

func NewBasePattern(pointCtor func() point.Point, width, height int) *BasePattern {
	newBase := &BasePattern{adjacentEdge: make(map[Orientation][]point.Point)}
	newBase.pointMatrix = make([][]point.Point, width)
	for x := 0; x < width; x++ {
		newBase.pointMatrix[x] = make([]point.Point, height)
		for y := 0; y < height; y++ {
			newBase.pointMatrix[x][y] = pointCtor()
		}
	}
	return newBase
}

func (base *BasePattern) GetDimensions() (int, int) {
	if len(base.pointMatrix) == 0 {
		return 0, 0
	}
	if len(base.pointMatrix[0]) == 0 {
		return len(base.pointMatrix), 0
	}
	return len(base.pointMatrix), len(base.pointMatrix[0])
}

func (base *BasePattern) Transform(transform Transformation) Pattern {
	return nil
}

func (base *BasePattern) GetEdge(orientation Orientation) []point.Point {
	var retSlice []point.Point
	switch orientation {
	case Left:
		for i := 0; i < len(base.pointMatrix); i++ {
			retSlice = append(retSlice, base.PointAt(0, i))
		}
	case Right:
		for i := 0; i < len(base.pointMatrix); i++ {
			retSlice = append(retSlice, base.PointAt(len(base.pointMatrix)-1, i))
		}
	case Top:
		for i := 0; i < len(base.pointMatrix); i++ {
			retSlice = append(retSlice, base.PointAt(i, 0))
		}
	case Bottom:
		for i := 0; i < len(base.pointMatrix); i++ {
			retSlice = append(retSlice, base.PointAt(i, len(base.pointMatrix)-1))
		}
	}
	return retSlice
}

func printHorizontal(pts []point.Point) {
	for i := 0; i < len(pts); i++ {
		point.PrintPoint(pts[i])
	}
	fmt.Println()
}

func printVertical(pts []point.Point) {
	for i := 0; i < len(pts); i++ {
		point.PrintPoint(pts[i])
		fmt.Println()
	}
	fmt.Println("###")
}

func (base *BasePattern) Accept(pattern Pattern, pOrientation Orientation) bool {
	var mySide, theirSide []point.Point
	myWidth, myHeight := base.GetDimensions()
	theirWidth, theirHeight := pattern.GetDimensions()
	mySide = base.GetEdge(pOrientation)
	switch pOrientation {
	case Top:
		if myWidth != theirWidth {
			return false
		}
		theirSide = pattern.GetEdge(Bottom)
	case Right:
		if myHeight != theirHeight {
			return false
		}
		theirSide = pattern.GetEdge(Left)
	case Bottom:
		if myWidth != theirWidth {
			return false
		}
		theirSide = pattern.GetEdge(Top)
	case Left:
		if myHeight != theirHeight {
			return false
		}
		theirSide = pattern.GetEdge(Right)
	}
	if len(mySide) != len(theirSide) {
		return false
	}
	for i := 0; i < len(mySide); i++ {
		if !mySide[i].Equal(theirSide[i]) {
			return false
		}
	}
	return true
}

func (base *BasePattern) SetSize(width, height int) {
	base.pointMatrix = make([][]point.Point, height)
	for i := 0; i < len(base.pointMatrix); i++ {
		base.pointMatrix[i] = make([]point.Point, width)
	}
}

func (base *BasePattern) PointAt(x, y int) point.Point {
	width, height := base.GetDimensions()
	if x > width || y > height {
		return nil
	}
	return base.pointMatrix[x][y]
}

func (base *BasePattern) SetPointAt(x, y int, pt point.Point) {
	base.pointMatrix[x][y] = pt
}

func (base *BasePattern) Equal(pat Pattern) bool {
	bw, bh := base.GetDimensions()
	pw, ph := base.GetDimensions()
	if bw != pw || bh != ph {
		return false
	}
	for x := 0; x < bw; x++ {
		for y := 0; y < bh; y++ {
			if !base.PointAt(x, y).Equal(pat.PointAt(x, y)) {
				return false
			}
		}
	}
	return true
}

func (base *BasePattern) SetAcceptingEdge(orientation Orientation, points ...point.Point) {
	base.adjacentEdge[orientation] = append(points)
}

func (base *BasePattern) GetAcceptingEdge(orientation Orientation) []point.Point {
	return base.adjacentEdge[orientation]
}
