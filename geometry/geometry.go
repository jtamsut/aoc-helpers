package geometry

import helpers "github.com/jtamsut/aochelpers"

// Box ... represents the coordinates of a box
type Box struct {
	LeftX, TopY, Width, Height int
}

// DetermineOverlap ... determines if there is an overlap between two boxes
func DetermineOverlap(p Box, m Box) (bool, Box) {
	lX := helpers.Max(p.LeftX, m.LeftX)
	rX := helpers.Min(p.LeftX+p.Width, m.LeftX+m.Width)
	bY := helpers.Max(p.TopY-p.Height, m.TopY-m.Height)
	tY := helpers.Min(p.TopY, m.TopY)
	var boxWidth, boxHeight int

	if rX > lX {
		boxWidth = rX - lX
	}

	if tY > bY {
		boxHeight = tY - bY
	}

	// return true, empty box if no overlap
	if boxHeight*boxWidth == 0 {
		return false, Box{}
	}

	// if overlap, return false and resulting overlapping box
	return true, Box{
		LeftX:  lX,
		TopY:   tY,
		Width:  boxWidth,
		Height: boxHeight,
	}
}

// Area ... gets area of a box
func Area(b Box) int {
	return b.Height * b.Width
}
