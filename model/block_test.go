package model

import (
	"reflect"
	"testing"
)

func TestBottomRow(t *testing.T) {
	for _, test := range []struct {
		desc          string
		block         *Block
		wantBottomRow [][2]uint8
	}{
		{
			desc: "facing up",
			block: &Block{
				X:     3,
				Y:     4,
				Kind:  TileL,
				Angle: DirectionUp,
			},
			wantBottomRow: [][2]uint8{
				{3, 6},
				{4, 6},
			},
		},
	} {
		t.Run(test.desc, func(t *testing.T) {
			gotBottomRow := bottomRow(test.block)
			if !reflect.DeepEqual(gotBottomRow, test.wantBottomRow) {
				t.Errorf("Got row %v, want %v", gotBottomRow, test.wantBottomRow)
			}
		})
	}
}
