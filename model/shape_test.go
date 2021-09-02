package model

import (
	"reflect"
	"testing"
)

func TestRotateT(t *testing.T) {
	for _, test := range []struct {
		desc      string
		direction Direction
		wantShape Shape
	}{
		{
			desc:      "rotate up",
			direction: DirectionUp,
			wantShape: shapes[TileT],
		},
		/*{
			desc: "rotate down",
			direction: DirectionDown,
		},
		{
			desc: "rotate left",
			direction: DirectionLeft,
		},
		{
			desc: "rotate right",
			direction: DirectionRight,
		},*/
	} {
		t.Run(test.desc, func(t *testing.T) {
			gotShape := rotate(shapes[TileT], test.direction)

			if !reflect.DeepEqual(gotShape, test.wantShape) {
				t.Errorf("shapes are not equal: %v (%T) vs. %v (%T)", gotShape, gotShape, test.wantShape, test.wantShape)
			}
		})
	}
}
