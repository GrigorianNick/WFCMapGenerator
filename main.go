package main

import (
	"github.com/grigoriannick/WFCMapGenerator/tile"

	"github.com/grigoriannick/WFCMapGenerator/pattern"
)

func main() {
	patterns := pattern.ImageToPatterns(`.\scratch\BigPattern.png`, 2, 2)
	ts := tile.NewTileSet(120, 30, patterns...)
	ts.Collapse()
	tile.PrintTileSet(ts)
}
