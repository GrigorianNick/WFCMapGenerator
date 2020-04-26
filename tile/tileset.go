package tile

import (
	"fmt"
	"log"

	"github.com/grigoriannick/WFCMapGenerator/pattern"
	"github.com/grigoriannick/WFCMapGenerator/point"
)

type TileSet struct {
	Tiles                 [][]Tile
	tileWidth, tileHeight int
}

type tileAddress struct {
	x, y int
}

func NewTileSet(width, height int, patternPool ...pattern.Pattern) *TileSet {
	newTileSet := &TileSet{tileWidth: width, tileHeight: height}
	newTileSet.Tiles = make([][]Tile, newTileSet.tileWidth)
	for i := 0; i < newTileSet.tileWidth; i++ {
		newTileSet.Tiles[i] = make([]Tile, newTileSet.tileHeight)
		for j := 0; j < newTileSet.tileHeight; j++ {
			newTileSet.Tiles[i][j].SetPossibilePatterns(patternPool...)
		}
	}
	return newTileSet
}

func (tileSet *TileSet) GetDimensions() (int, int) {
	return tileSet.tileWidth, tileSet.tileHeight
}

func (tileSet *TileSet) GetCollapseOrder() []tileAddress {
	var retSlice []tileAddress
	for x := range tileSet.Tiles {
		for y := range tileSet.Tiles[x] {
			retSlice = append(retSlice, tileAddress{x: x, y: y})
		}
	}
	var tmp []tileAddress
	for _, addr := range retSlice {
		tmp = append([]tileAddress{addr}, tmp...)
	}
	return tmp
}

func (tileSet *TileSet) Collapse() {
	toCollapse := tileSet.GetCollapseOrder()
	var history []tileAddress
	for len(history) < len(toCollapse) {
		nextTile := len(history)
		neighbors := []neighbor{
			neighbor{
				tile:     tileSet.GetTile(toCollapse[nextTile].x-1, toCollapse[nextTile].y),
				relation: pattern.Left},
			neighbor{
				tile:     tileSet.GetTile(toCollapse[nextTile].x+1, toCollapse[nextTile].y),
				relation: pattern.Right},
			neighbor{
				tile:     tileSet.GetTile(toCollapse[nextTile].x, toCollapse[nextTile].y-1),
				relation: pattern.Top},
			neighbor{
				tile:     tileSet.GetTile(toCollapse[nextTile].x, toCollapse[nextTile].y+1),
				relation: pattern.Bottom}}
		if !tileSet.Tiles[toCollapse[nextTile].x][toCollapse[nextTile].y].Reselect(neighbors...) {
			tileSet.Tiles[toCollapse[nextTile].x][toCollapse[nextTile].y].Reset()
			if len(history) > 1 {
				history = history[:len(history)-1]
			} else {
				log.Println("No valid tilesets.")
				return
			}
		} else {
			history = append(history, toCollapse[nextTile])
		}
	}
}

func (tileSet *TileSet) GetTile(x, y int) *Tile {
	if x < 0 || y < 0 || x >= tileSet.tileWidth || y >= tileSet.tileHeight {
		return nil
	}

	return &tileSet.Tiles[x][y]
}

func PrintTileSet(tileSet *TileSet) {
	setWidth, setHeight := tileSet.GetDimensions()
	tileWidth, tileHeight := tileSet.Tiles[0][0].patternPool[0].GetDimensions()
	log.Println(tileSet.GetDimensions())
	for sy := 0; sy < setHeight; sy++ {
		for y := 0; y < tileHeight; y++ {
			for sx := 0; sx < setWidth; sx++ {
				for x := 0; x < tileWidth; x++ {
					point.PrintPoint(tileSet.GetTile(sx, sy).GetSelectedPattern().PointAt(x, y))
				}
				//fmt.Print("-")
			}
			fmt.Println()
		}
		//fmt.Println("|||| |||| ||||")
	}
}
