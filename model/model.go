package model

import "time"

type TileKind uint8

const (
	TileS TileKind = iota
	TileZ
	TileL
	TileJ
	TileI
	TileO
	TileT
)

type Direction uint8

const (
	DirectionUp Direction = iota
	DirectionDown
	DirectionLeft
	DirectionRight
)

type Block struct {
	X     uint8
	Y     uint8
	Angle uint8 // 0-3; represents the multiple 90 degrees.
	Kind  TileKind
}

type Model interface {
	Data() [][]TileKind
	CurrentBlock() Block
	NextBlock() Block
	Score() uint32
	Level() uint8

	Update(dt time.Duration)
	MoveCurrentBlock(direction Direction)
}

type model struct {
}

func New() Model {
	return &model{}
}

func (m *model) Data() [][]TileKind  { return nil }
func (m *model) CurrentBlock() Block { return Block{} }
func (m *model) NextBlock() Block    { return Block{} }
func (m *model) Score() uint32       { return 0 }
func (m *model) Level() uint8        { return 0 }

func (m *model) Update(dt time.Duration)              {}
func (m *model) MoveCurrentBlock(direction Direction) {}
