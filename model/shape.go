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
				result[x][y] = s[ShapeWidth-x-1][ShapeHeight-y-1]
			}
		}

		return trimTopRow(result)
	case DirectionLeft:
		result := Shape{}
		for x := 0; x < ShapeHeight; x++ {
			result = append(result, []bool{false, false, false, false})
			for y := 0; y < ShapeWidth; y++ {
				result[x][y] = s[ShapeWidth-y-1][ShapeHeight-x-1]
			}
		}

		return trimFirstColumn(result)
	case DirectionRight:
		result := Shape{}
		for x := 0; x < ShapeHeight; x++ {
			result = append(result, []bool{false, false, false, false})
			for y := 0; y < ShapeWidth; y++ {
				result[x][y] = s[y][x]
			}
		}

		return trimFirstColumn(result)
	}

	return nil
}

func trimTopRow(s Shape) Shape {
	for _, isSet := range s[0] {
		if isSet {
			return s
		}
	}

	topRow := s[0]
	s = s[1:]
	s = append(s, topRow)
	return s
}

func trimFirstColumn(s Shape) Shape {
	for _, row := range s {
		if row[0] {
			return s
		}
	}

	for idx := range s {
		s[idx] = s[idx][1:]
		s[idx] = append(s[idx], false)
	}

	return s
}
