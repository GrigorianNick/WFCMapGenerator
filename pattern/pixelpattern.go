package pattern

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"os"

	"github.com/grigoriannick/WFCMapGenerator/point"
)

func PrintPatten(pattern Pattern) {
	pixPattern, ok := pattern.(*PixelPattern)
	if !ok {
		return
	}
	w, h := pixPattern.GetDimensions()
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			point.PrintPoint(pixPattern.PointAt(x, y))
		}
		fmt.Println("-")
	}
	/*fmt.Print("Top   :")
	printHorizontal(pattern.GetEdge(Top))
	fmt.Print("Bottom:")
	printHorizontal(pattern.GetEdge(Bottom))
	fmt.Print("Left  :")
	printHorizontal(pattern.GetEdge(Left))
	fmt.Print("Right :")
	printHorizontal(pattern.GetEdge(Right))*/
	fmt.Println("+")
}

func NewPixelPattern(width, height int) *PixelPattern {
	return &PixelPattern{
		BasePattern: NewBasePattern(func() point.Point { return point.NewPixelPoint() }, width, height)}
}

func ImageToPatterns(path string, patternWidth, patternHeight int) []Pattern {
	var retSlice []Pattern
	file, err := os.Open(path)
	if err != nil {
		log.Println(err)
		return retSlice
	}
	imgInfo, _, err := image.DecodeConfig(file)
	if err != nil {
		log.Println(err)
		return retSlice
	}
	file.Seek(0, 0)

	img, _, err := image.Decode(file)
	if err != nil {
		log.Println(err)
		return retSlice
	}

	numHPatterns := imgInfo.Width / patternWidth
	if imgInfo.Width%patternWidth != 0 {
		numHPatterns++
	}

	numVPatterns := imgInfo.Height / patternHeight
	if imgInfo.Height%patternHeight != 0 {
		numVPatterns++
	}

	var patternMatrix [][]Pattern

	patternMatrix = make([][]Pattern, numHPatterns)

	for i := 0; i < numHPatterns; i++ {
		patternMatrix[i] = make([]Pattern, numVPatterns)
		for j := 0; j < numVPatterns; j++ {
			patternMatrix[i][j] = NewPixelPattern(patternWidth, patternHeight)
		}
	}

	// Fill out the point matricies
	for x := 0; x < imgInfo.Width; x++ {
		for y := 0; y < imgInfo.Height; y++ {
			r, g, b, a := img.At(x, y).RGBA()
			//RGBA() premultiplies by alpha, so we need to reverse that
			r = r / a
			g = g / a
			b = b / a
			patternMatrix[x/patternWidth][y/patternHeight].SetPointAt(
				x%patternWidth,
				y%patternHeight,
				&point.PixelPoint{
					Pixel: color.RGBA{
						R: uint8(r),
						G: uint8(g),
						B: uint8(b),
						A: uint8(a)}})
		}
	}

	// Build accepting edges
	for x, col := range patternMatrix {
		for y, pat := range col {
			if x > 0 {
				pat.SetAcceptingEdge(
					Left,
					patternMatrix[x-1][y].GetEdge(Right)...)
			}
			if x < len(patternMatrix)-1 {
				pat.SetAcceptingEdge(
					Right,
					patternMatrix[x+1][y].GetEdge(Left)...)
			}
			if y > 0 {
				pat.SetAcceptingEdge(
					Top,
					patternMatrix[x][y-1].GetEdge(Bottom)...)
			}
			if y < len(col)-1 {
				pat.SetAcceptingEdge(
					Bottom,
					patternMatrix[x][y+1].GetEdge(Top)...)
			}
		}
	}

	for _, x := range patternMatrix {
		retSlice = append(retSlice, x...)
	}
	return retSlice
}

type PixelPattern struct {
	*BasePattern
}
