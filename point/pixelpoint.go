package point

import (
	"fmt"
	"image/color"
)

type PixelPoint struct {
	Pixel color.RGBA
}

func NewPixelPoint() *PixelPoint {
	return &PixelPoint{}
}

func (pixelPoint *PixelPoint) Equal(pattern Point) bool {
	pix, ok := pattern.(*PixelPoint)
	if !ok {
		return false
	}
	rr, rg, rb, ra := pixelPoint.Pixel.RGBA()
	pr, pg, pb, pa := pix.Pixel.RGBA()
	return rr == pr && rg == pg && rb == pb && ra == pa
}

func PrintPoint(pt Point) {
	pix, ok := pt.(*PixelPoint)
	if !ok {
		return
	}
	r, _, _, _ := pix.Pixel.RGBA()
	if r > 0 {
		fmt.Print(" ")
	} else {
		fmt.Print("X")
	}
}
