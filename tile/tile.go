package tile

import (
	"math/rand"

	"github.com/grigoriannick/WFCMapGenerator/pattern"
)

type Tile struct {
	patternPool        []pattern.Pattern
	availablePatterns  []pattern.Pattern
	selectedPattern    pattern.Pattern
	transformationPool []pattern.Transformation
}

type neighbor struct {
	tile     *Tile
	relation pattern.Orientation
}

func (tile *Tile) GetEntropy() int {
	return len(tile.patternPool)
}

// Filter available patterns based on neighbors
func (tile *Tile) prune(neighbors ...neighbor) {
	for _, n := range neighbors {
		if n.tile == nil || n.tile.selectedPattern == nil {
			continue
		}
		var acceptedPatterns []pattern.Pattern
		for _, p := range tile.availablePatterns {
			if p.Accept(n.tile.selectedPattern, n.relation) {
				acceptedPatterns = append(acceptedPatterns, p)
			}
		}
		tile.availablePatterns = acceptedPatterns
	}
}

func (tile *Tile) collapse(neighbors ...neighbor) bool {
	tile.prune(neighbors...)
	if len(tile.availablePatterns) == 0 {
		return false
	}
	tile.selectedPattern = tile.availablePatterns[rand.Intn(len(tile.availablePatterns))]
	return true
}

func (tile *Tile) Reset() {
	tile.availablePatterns = append([]pattern.Pattern(nil), tile.patternPool...)
	tile.selectedPattern = nil
}

// Reselect removes the tile we've selected from our available pool, and tries again
func (tile *Tile) Reselect(neighbors ...neighbor) bool {
	// If we picked a tile previously, remove it from our available pool
	if tile.selectedPattern != nil {
		var newAvailablePatterns []pattern.Pattern
		for _, p := range tile.availablePatterns {
			if !tile.selectedPattern.Equal(p) {
				newAvailablePatterns = append(newAvailablePatterns, p)
			}
		}
		tile.availablePatterns = newAvailablePatterns
		tile.selectedPattern = nil
	}
	return tile.collapse(neighbors...)
}

func (tile *Tile) SetPossibilePatterns(patterns ...pattern.Pattern) {
	tile.patternPool = []pattern.Pattern{}
	tile.selectedPattern = nil
	for _, pattern := range patterns {
		tile.patternPool = append(tile.patternPool, pattern)
		for _, transformation := range tile.transformationPool {
			tile.patternPool = append(tile.patternPool, pattern.Transform(transformation))
		}
	}
	tile.Reset()
}

func (tile *Tile) GetSelectedPattern() pattern.Pattern {
	return tile.selectedPattern
}
