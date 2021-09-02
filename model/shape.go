package model

type Shape [][]bool

var shapes = map[TileKind]Shape{
	TileS: {
		{true, false},
		{true, true},
		{false, true},
		{false, false},
	},
	TileZ: {
		{false, true},
		{true, true},
		{true, false},
		{false, false},
	},
	TileL: {
		{true, false},
		{true, false},
		{true, true},
		{false, false},
	},
	TileJ: {
		{false, true},
		{false, true},
		{true, true},
		{false, false},
	},
	TileI: {
		{true, false},
		{true, false},
		{true, false},
		{true, false},
	},
	TileO: {
		{true, true},
		{true, true},
		{false, false},
		{false, false},
	},
	TileT: {
		{true, false},
		{true, true},
		{true, false},
		{false, false},
	},
}

func rotate(s Shape, d Direction) Shape {
	switch d {
	case DirectionUp:
		return s
	case DirectionDown:
		result := Shape{}
		for x := 0; x < ShapeWidth; x++ {
			result = append(result, []bool{false, false})
			for y := 0; y < ShapeHeight; y++ {
				result[ShapeWidth-x-1][ShapeHeight-y-1] = s[x][y]
			}
		}
		return result
	case DirectionLeft:
	case DirectionRight:
	}

	return nil
}
