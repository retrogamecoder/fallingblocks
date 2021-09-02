package model

type Block struct {
	X     uint8
	Y     uint8
	Angle Direction
	Kind  TileKind
}

func bottomRow(b *Block) [][2]uint8 {
	rotatedShape := rotate(shapes[b.Kind], b.Angle)

	lowestY := make([]uint8, len(rotatedShape[0]))

	for y, row := range rotatedShape {
		for x, isSet := range row {
			if !isSet {
				continue
			}

			lowestY[x] = b.Y - uint8(y)
		}
	}

	return nil
}
